package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Walker088/rigel_ledger_server/src/golang/config"
)

type PgPool struct {
	pool *pgxpool.Pool
}

func New(config *config.PgConfig) (*PgPool, error) {
	poolConf, err := pgxpool.ParseConfig(config.ToConnString())
	if err != nil {
		return nil, err
	}
	poolConf.MinConns = config.MinConns
	poolConf.MaxConns = config.MaxConns
	poolConf.MaxConnIdleTime = config.MaxConnIdleTime
	poolConf.MaxConnLifetime = config.MaxConnLifetime
	poolConf.MaxConnLifetimeJitter = config.MaxConnLifetimeJitter
	poolConf.HealthCheckPeriod = config.HealthCheckPeriod

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConf)
	if err != nil {
		return nil, err
	}
	return &PgPool{pool: pool}, nil
}

func (p *PgPool) ShutDownPool() {
	p.pool.Close()
}

func (p *PgPool) GetPool() *pgxpool.Pool {
	return p.pool
}
