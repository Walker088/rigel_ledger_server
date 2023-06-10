package backend

import (
	"github.com/Walker088/rigel_ledger_server/src/golang/jwt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Context struct {
	Pool   *pgxpool.Pool
	Jwt    *jwt.JwtEngine
	Logger *zap.SugaredLogger
}

func New(p *pgxpool.Pool, jwt *jwt.JwtEngine, l *zap.SugaredLogger) *Context {
	return &Context{
		Pool:   p,
		Jwt:    jwt,
		Logger: l,
	}
}
