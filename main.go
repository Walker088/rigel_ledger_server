package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Walker088/rigel_ledger_server/backend/config"
	db "github.com/Walker088/rigel_ledger_server/backend/database"
	log "github.com/Walker088/rigel_ledger_server/backend/logger"
	"github.com/Walker088/rigel_ledger_server/backend/router"
)

func main() {
	c := config.GetAppConfig()
	l := log.New()
	defer l.Sync()

	pool := db.PgPool{Logger: l}
	pool.StartPool()
	defer pool.GetPool().Close()

	l.Info("Welcome to Rigel Ledger")
	r := router.New(&c.GithubOAuthConfig, l)
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", c.AppHost, c.AppPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	l.Fatal(srv.ListenAndServe())
}
