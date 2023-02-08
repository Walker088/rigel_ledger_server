package middlewares

import (
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"
)

type MiddleWares struct {
	logger       *zap.SugaredLogger
	AllowMethods []string
	AllowOrigins []string
}

func New(logger *zap.SugaredLogger, origins []string) *MiddleWares {
	return &MiddleWares{
		logger:       logger,
		AllowOrigins: origins,
		AllowMethods: nil,
	}
}

func (m *MiddleWares) AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			m.logger.Infof("[From: %s] [To: %s] tooks %s",
				r.RemoteAddr,
				r.RequestURI,
				time.Since(start),
			)
		}()
		next.ServeHTTP(w, r)
	})
}

func (m *MiddleWares) DefaultRestHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rqOrigin := r.Header.Get("Origin")
		for _, h := range m.AllowOrigins {
			if h == rqOrigin {
				w.Header().Set("Access-Control-Allow-Origin", rqOrigin)
			}
		}
		if m.AllowMethods != nil {
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(m.AllowMethods, ","))
		}
		w.Header().Set("Content-Type", "application/json")
		//w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Date", time.Now().Format(time.RFC1123))
		next.ServeHTTP(w, r)
	})
}

func (m *MiddleWares) ValidateJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}
