package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Walker088/rigel_ledger_server/backend"
	"github.com/Walker088/rigel_ledger_server/backend/config"
	"github.com/Walker088/rigel_ledger_server/backend/database"
	"github.com/Walker088/rigel_ledger_server/backend/jwt"
	"github.com/Walker088/rigel_ledger_server/backend/logger"
	"github.com/Walker088/rigel_ledger_server/backend/router"
)

func main() {
	c := config.GetAppConfig()
	l := logger.New()
	defer l.Sync()

	pool, err := database.New(config.GetPgConfig())
	if err != nil {
		l.DPanicf("Init DB Conn Pool error: %w", err)
	}
	defer pool.ShutDownPool()

	jwt := jwt.New([]byte(c.JwtSecret), l)
	gctx := backend.New(pool.GetPool(), jwt, l)

	l.Info("Welcome to RigelLedger")
	m := router.New(c, gctx)
	srv := &http.Server{
		Handler:      m.Router,
		Addr:         fmt.Sprintf("%s:%s", c.AppHost, c.AppPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	go func() {
		l.Error(srv.ListenAndServe())
	}()

	deadlineChannel := make(chan os.Signal, 1)
	signal.Notify(deadlineChannel, os.Interrupt)
	<-deadlineChannel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	l.Info("Shutting down RigelLedger REST Server")
	os.Exit(0)
}
