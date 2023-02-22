package dao

import "github.com/jackc/pgx/v5/pgxpool"

type LedgerJournalDao struct {
	pool *pgxpool.Pool
}

func NewLedgerJournalDao(p *pgxpool.Pool) *LedgerJournalDao {
	return &LedgerJournalDao{
		pool: p,
	}
}

func (l *LedgerJournalDao) Get(userId string, transacId int) {}

func (l *LedgerJournalDao) GetList(userId string, limit int, offset int) {}

func (l *LedgerJournalDao) Create() {}

func (l *LedgerJournalDao) Update(userId string, transacId int) {}
