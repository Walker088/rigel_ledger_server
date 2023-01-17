package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type MiddleWares struct {
	logger       *zap.SugaredLogger
	AllowMethods []string
}

func New(logger *zap.SugaredLogger) *MiddleWares {
	return &MiddleWares{
		logger:       logger,
		AllowMethods: nil,
	}
}

func (m *MiddleWares) AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		defer func() {
			m.logger.Infof("[%d] [From: %s] [To: %s] tooks %s [%d bytes written]",
				ww.Status(),
				r.RemoteAddr,
				r.RequestURI,
				time.Since(start),
				ww.BytesWritten(),
			)
		}()
		next.ServeHTTP(w, r)
	})
}

func (m *MiddleWares) DefaultRestHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.AllowMethods != nil {
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(m.AllowMethods, ","))
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Date", time.Now().Format(time.RFC1123))
		next.ServeHTTP(w, r)
	})
}
