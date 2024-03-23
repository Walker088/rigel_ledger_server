package main

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	backend "github.com/Walker088/rigel_ledger_server/src/golang"
	"github.com/Walker088/rigel_ledger_server/src/golang/config"
	"github.com/Walker088/rigel_ledger_server/src/golang/database"
	"github.com/Walker088/rigel_ledger_server/src/golang/jwt"
	"github.com/Walker088/rigel_ledger_server/src/golang/logger"
	"github.com/Walker088/rigel_ledger_server/src/golang/router"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func migrate(pool *pgxpool.Pool) {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	db := stdlib.OpenDBFromPool(pool)

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func main() {
	c := config.GetAppConfig()
	l := logger.New()
	defer l.Sync()

	pool, err := database.New(config.GetPgConfig())
	if err != nil {
		l.DPanicf("Init DB Conn Pool error: %w", err)
	}
	defer pool.ShutDownPool()

	migrate(pool.GetPool())
	jwt := jwt.New([]byte(c.JwtSecret), l)
	gctx := backend.New(pool.GetPool(), jwt, l)

	l.Info("Welcome to RigelLedger")
	m := router.New(c, gctx)
	addr := fmt.Sprintf("%s:%s", c.AppHost, c.AppPort)
	srv := &http.Server{
		Handler:      m.Router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	go func() {
		l.Infof("Server started at %s", addr)
		l.Warn(srv.ListenAndServe())
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
