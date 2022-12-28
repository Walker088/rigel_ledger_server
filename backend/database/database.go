package database

import "github.com/jackc/pgx/v5/pgxpool"

type DbConnPool interface {
	ShutDownPool()
	StartPool() error
	GetPool() *pgxpool.Pool
}
