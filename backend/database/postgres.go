package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/Walker088/rigel_ledger_server/backend/config"
)

type PgPool struct {
	Logger *zap.SugaredLogger
	pool   *pgxpool.Pool
}

func (p *PgPool) ShutDownPool() {
	p.pool.Close()
}

func (p *PgPool) StartPool() error {
	var err error
	var poolConf *pgxpool.Config
	var pool *pgxpool.Pool

	appPgConf := config.GetPgConfig()
	poolConf, err = pgxpool.ParseConfig(appPgConf.ToConnString())
	if err != nil {
		p.Logger.Errorf("Unable to parse DATABASE_URL: %v", err)
		return err
	}
	poolConf.MinConns = appPgConf.MinConns
	poolConf.MaxConns = appPgConf.MaxConns
	poolConf.MaxConnIdleTime = appPgConf.MaxConnIdleTime
	poolConf.MaxConnLifetime = appPgConf.MaxConnLifetime
	poolConf.MaxConnLifetimeJitter = appPgConf.MaxConnLifetimeJitter
	poolConf.HealthCheckPeriod = appPgConf.HealthCheckPeriod

	pool, err = pgxpool.NewWithConfig(context.Background(), poolConf)
	if err != nil {
		p.Logger.Errorf("Unable to create connection pool: %v", err)
		return err
	}
	p.pool = pool
	return nil
}

func (p *PgPool) GetPool() *pgxpool.Pool {
	return p.pool
}
