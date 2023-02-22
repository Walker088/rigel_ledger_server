package dao

import "github.com/jackc/pgx/v5/pgxpool"

type LedgerDao struct {
	pool *pgxpool.Pool
}

func NewLedgerDao(p *pgxpool.Pool) *LedgerDao {
	return &LedgerDao{
		pool: p,
	}
}

func (l *LedgerDao) Get(userId string) {}

func (l *LedgerDao) GetList(userId string) {}

func (l *LedgerDao) Create() {}

func (l *LedgerDao) Update(userId string, transacId int) {}
