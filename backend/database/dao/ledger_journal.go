package dao

import (
	"context"

	"github.com/Walker088/rigel_ledger_server/backend/database/pojo"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LedgerJournalDao struct {
	pool *pgxpool.Pool
}

func NewLedgerJournalDao(p *pgxpool.Pool) *LedgerJournalDao {
	return &LedgerJournalDao{
		pool: p,
	}
}

func (l *LedgerJournalDao) Get(userId string, transacId int) (*pojo.LedgerJournal, error) {
	var ledger *pojo.LedgerJournal
	queryLedgerInfo := `
	
	`

	if err := pgxscan.Get(
		context.Background(),
		l.pool,
		&ledger,
		queryLedgerInfo,
		userId,
		transacId,
	); err != nil {
		return ledger, err
	}
	return ledger, nil
}

func (l *LedgerJournalDao) GetList(userId string, limit int, offset int) ([]*pojo.LedgerJournal, error) {
	var ledger []*pojo.LedgerJournal
	queryLedgerInfo := `
	
	`
	if err := pgxscan.Select(
		context.Background(),
		l.pool,
		&ledger,
		queryLedgerInfo,
		userId,
		limit,
		offset,
	); err != nil {
		return nil, err
	}

	return ledger, nil
}

func (l *LedgerJournalDao) Create(entries []*pojo.CreateJournalEntry) ([]*pojo.CreateJournalEntry, error) {
	var createdEntries = make([]*pojo.CreateJournalEntry, 0)
	insertJournalQuery := ``
	tx, err := l.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	for _, entry := range entries {
		var created *pojo.CreateJournalEntry
		if err := pgxscan.Get(
			context.Background(),
			tx,
			created,
			insertJournalQuery,
			entry.UserId,
		); err != nil {
			return nil, err
		}
		createdEntries = append(createdEntries, created)
	}

	return createdEntries, nil
}

func (l *LedgerJournalDao) RevertTx(userId string, transacId int) (*pojo.CreateJournalEntry, error) {
	var ajustedTx *pojo.CreateJournalEntry
	var originalTx *pojo.LedgerJournal
	queryTransac := ``
	if err := pgxscan.Get(context.Background(), l.pool, &originalTx, queryTransac, transacId); err != nil {
		return ajustedTx, err
	}

	revertTransacQuery := ``
	if err := pgxscan.Get(
		context.Background(),
		l.pool,
		ajustedTx,
		revertTransacQuery,
	); err != nil {
		return ajustedTx, err
	}

	return ajustedTx, nil
}
