package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chimdw "github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	backend "github.com/Walker088/rigel_ledger_server/src/golang"
	"github.com/Walker088/rigel_ledger_server/src/golang/config"
	custommdw "github.com/Walker088/rigel_ledger_server/src/golang/router/middlewares"
	"github.com/Walker088/rigel_ledger_server/src/golang/router/v1/protected"
	"github.com/Walker088/rigel_ledger_server/src/golang/router/v1/public"
	"github.com/Walker088/rigel_ledger_server/src/golang/router/v1/public/oauth"
)

type Mux struct {
	Router *chi.Mux
	logger *zap.SugaredLogger
}

func (m *Mux) initRoutes(c *config.AppConfig, mw *custommdw.MiddleWares, gctx *backend.Context) {

	m.Router.Route("/v1/public", func(r chi.Router) {
		auth := oauth.New(c.OauthGithubClientId, c.OauthGithubClientSecret, m.logger, gctx.Jwt)
		home := public.NewHomeInfo()
		r.Get("/home", home.HomeInfoHandler)

		r.Route("/oauth/github", func(r chi.Router) {
			r.Get("/login", auth.GithubLogin)
		})
	})
	m.Router.Route("/v1/protected/user", func(r chi.Router) {
		r.Use(mw.ValidateJwt)

		uh := protected.NewUserHandler(gctx.Pool)
		r.Get("/{userId}/basic", uh.GetUserBasicHandler)
		r.Get("/{userId}/complete", uh.GetUserCompleteHandler)
		r.Patch("/{userId}/update", uh.UpdateUserHandler)

		lh := protected.NewLedgerHandler(gctx.Pool)
		r.Get("/{userId}/ledgers", lh.GetLedgerLstHandler)
		r.Get("/{userId}/ledgers/{ledgerId}", lh.GetLedgerHandler)
		r.Post("/{userId}/ledgers", lh.CreateLedgerHandler)
		r.Patch("/{userId}/ledgers/{ledgerId}", lh.UpdateLedgerHandler)

		jh := protected.NewLedgerJournalHandler(gctx.Pool)
		r.Get("/{userId}/journal", jh.GetJournalLstHandler)
		r.Get("/{userId}/journal/{transacId}", jh.GetJournalHandler)
		r.Post("/{userId}/journal", jh.CreateJournalHandler)
		r.Patch("/{userId}/journal/{transacId}", jh.RevertJournalHandler)

		rh := protected.NewReportHandler(gctx.Pool)
		r.Get("/{userId}/report/balancesheet/annual/year/{year}", rh.GetAnnualBalanceSheet)
		r.Get("/{userId}/report/balancesheet/seasonal/year/{year}/season/{season}", rh.GetSeasonalBalanceSheet)
		r.Get("/{userId}/report/incomestatement/annual/year/{year}", rh.GetAnnualIncomeStatement)
		r.Get("/{userId}/report/incomestatement/seasonal/year/{year}/season/{season}", rh.GetSeasonalIncomeStatement)
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

func New(c *config.AppConfig, gctx *backend.Context) *Mux {
	//compressor := chimdw.NewCompressor(4)

	r := chi.NewRouter()
	mw := custommdw.New(gctx.Logger, c.AppAllowOrigins)
	r.Use(chimdw.RealIP)
	//r.Use(compressor.Handler)
	r.Use(mw.DefaultRestHeaders)
	r.Use(chimdw.Recoverer)
	r.Use(mw.AccessLog)

	m := &Mux{
		Router: r,
		logger: gctx.Logger,
	}
	m.initRoutes(c, mw, gctx)
	mw.AllowMethods = m.getChiRouteMethods()

	m.logger.Info("Chi Mux Initialized")
	return m
}
