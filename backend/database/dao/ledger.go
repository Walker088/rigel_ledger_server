package dao

import (
	"github.com/Walker088/rigel_ledger_server/backend/database/pojo"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LedgerDao struct {
	pool *pgxpool.Pool
}

func NewLedgerDao(p *pgxpool.Pool) *LedgerDao {
	return &LedgerDao{
		pool: p,
	}
}

func (l *LedgerDao) Get(userId string, ledger_id string) (*pojo.LedgerInfo, error) {
	var ledgerInfo pojo.LedgerInfo

	return &ledgerInfo, nil
}

func (l *LedgerDao) GetList(userId string) ([]*pojo.LedgerInfo, error) {
	var userLedgers []*pojo.LedgerInfo

	return userLedgers, nil
}

func (l *LedgerDao) Create(ledgerInfo *pojo.LedgerInfo) {}

func (l *LedgerDao) Update(userId string, updated *pojo.UpdateLedgerInfo) {}
