package router

import (
	//"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimdw "github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"github.com/Walker088/rigel_ledger_server/backend/config"
	"github.com/Walker088/rigel_ledger_server/backend/jwt"
	custommdw "github.com/Walker088/rigel_ledger_server/backend/router/middlewares"
	"github.com/Walker088/rigel_ledger_server/backend/router/v1/protect"
	"github.com/Walker088/rigel_ledger_server/backend/router/v1/public"
	"github.com/Walker088/rigel_ledger_server/backend/router/v1/public/oauth"
)

type Mux struct {
	Router    *chi.Mux
	logger    *zap.SugaredLogger
	jwtEngine *jwt.JwtEngine
}

func (m *Mux) initRoutes(c *config.AppConfig) {
	//var completeHost = fmt.Sprintf("http://%s:%s", c.AppHost, c.AppPort)
	var auth = oauth.New(c.OauthGithubClientId, c.OauthGithubClientSecret, m.logger, m.jwtEngine)
	var home = public.NewHomeInfo()

	m.Router.Route("/v1/public", func(r chi.Router) {
		r.Get("/home", home.HomeInfoHandler)

		r.Route("/oauth/github", func(r chi.Router) {
			r.Get("/login", auth.GithubLogin)
		})
	})
	m.Router.Route("/v1/protect", func(r chi.Router) {
		r.Get("/{userId}", protect.UserHomeHandler)
	})
}

func (m *Mux) getChiRouteMethods() []string {
	allowMethods := make(map[string]struct{})
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		allowMethods[method] = struct{}{}
		return nil
	}
	if err := chi.Walk(m.Router, walkFunc); err != nil {
		m.logger.Errorf("Logging err: %s", err.Error())
	}

	if _, ok := allowMethods[http.MethodOptions]; ok {
		methods := make([]string, 0, len(allowMethods))
		for method := range allowMethods {
			methods = append(methods, method)
		}
		return methods
	}
	return nil
}

func New(c *config.AppConfig, logger *zap.SugaredLogger, jwtEngine *jwt.JwtEngine) *Mux {
	//compressor := chimdw.NewCompressor(4)

	r := chi.NewRouter()
	mw := custommdw.New(logger, c.AppAllowOrigins)
	r.Use(chimdw.RealIP)
	//r.Use(compressor.Handler)
	r.Use(mw.DefaultRestHeaders)
	r.Use(chimdw.Recoverer)
	r.Use(mw.AccessLog)

	m := &Mux{
		Router:    r,
		logger:    logger,
		jwtEngine: jwtEngine,
	}
	m.initRoutes(c)
	mw.AllowMethods = m.getChiRouteMethods()

	m.logger.Info("Chi Mux Initialized")
	return m
}
