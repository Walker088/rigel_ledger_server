package dao

import (
	"context"
	"fmt"

	"github.com/Walker088/rigel_ledger_server/src/golang/database/pojo"
	"github.com/georgysavva/scany/v2/pgxscan"
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

func (l *LedgerDao) checkLedgerOwnership(userId string, ledgerId string) bool {
	isLedgerOwner := false
	query := `SELECT ledger_owner = $1 FROM user_ledgers WHERE ledger_id = $2`
	if err := pgxscan.Get(context.Background(), l.pool, &isLedgerOwner, query, userId, ledgerId); err != nil {
		return isLedgerOwner
	}
	return isLedgerOwner
}

func (l *LedgerDao) Get(userId string, ledgerId string) (*pojo.LedgerInfo, error) {
	var ledgerInfo pojo.LedgerInfo

	isLedgerOwner := l.checkLedgerOwnership(userId, ledgerId)
	if !isLedgerOwner {
		return &ledgerInfo, fmt.Errorf("user %s is not owner of the ledger %s", userId, ledgerId)
	}

	queryLedger := `
	SELECT 
		l.ledger_id,
		l.ledger_owner,
		u.user_name ledger_owner_name,
		l.ledger_type_id,
		(SELECT type_name FROM ref_ledger_types WHERE ledger_type_id = l.ledger_type_id) ledger_type_name,
		l.currency,
		(SELECT currency_name FROM ref_currencies_iso4217 WHERE alphabetic_code = l.currency),
		l.balance,
		CASE
			WHEN u.main_currency = l.currency THEN l.balance
			ELSE l.balance 
				/ (SELECT exchange_rate FROM ref_exchange_rates WHERE from_code = 'USD' AND to_code = l.currency)
				* (SELECT exchange_rate FROM ref_exchange_rates WHERE from_code = 'USD' AND to_code = u.main_currency)
		END balance_adjusted,
		l.created_at,
		l.updated_at
	FROM 
		user_ledgers l
	JOIN user_info u ON (l.ledger_owner = u.user_id)
	WHERE
		l.ledger_id = $1
	`
	if err := pgxscan.Get(context.Background(), l.pool, &ledgerInfo, queryLedger, ledgerId); err != nil {
		return &ledgerInfo, err
	}
	return &ledgerInfo, nil
}

func (l *LedgerDao) GetList(userId string) ([]*pojo.LedgerInfo, error) {
	var userLedgers []*pojo.LedgerInfo
	queryLedgers := `
	SELECT 
		l.ledger_id,
		l.ledger_owner,
		u.user_name ledger_owner_name,
		l.ledger_type_id,
		(SELECT type_name FROM ref_ledger_types WHERE ledger_type_id = l.ledger_type_id) ledger_type_name,
		l.currency,
		(SELECT currency_name FROM ref_currencies_iso4217 WHERE alphabetic_code = l.currency),
		l.balance,
		CASE
			WHEN u.main_currency = l.currency THEN l.balance
			ELSE l.balance 
				/ (SELECT exchange_rate FROM ref_exchange_rates WHERE from_code = 'USD' AND to_code = l.currency)
				* (SELECT exchange_rate FROM ref_exchange_rates WHERE from_code = 'USD' AND to_code = u.main_currency)
		END balance_adjusted,
		l.created_at,
		l.updated_at
	FROM 
		user_ledgers l
	JOIN user_info u ON (l.ledger_owner = u.user_id)
	WHERE
		l.ledger_owner = $1`
	if err := pgxscan.Select(context.Background(), l.pool, &userLedgers, queryLedgers, userId); err != nil {
		return nil, err
	}
	return userLedgers, nil
}

func (l *LedgerDao) Create(ledgerInfo []*pojo.LedgerInfo) ([]*pojo.LedgerInfo, error) {
	var createdLedgers = make([]*pojo.LedgerInfo, 0)
	insertLedgerQuery := `
	INSERT INTO user_ledgers (ledger_id, ledger_owner, ledger_name, ledger_type_id, currency, balance)
	SELECT
		CONCAT($1, '_', $3, '_', COUNT(1) + 1,
		$1,
		$2,
		$3,
		$4,
		$5
	FROM
		user_ledgers
	WHERE
		ledger_owner = $1
		AND ledger_type_id = $3
	RETURNING (
		ledger_id,
		ledger_owner,
		ledger_name,
		ledger_type_id,
		currency,
		balance
	)
	`
	tx, err := l.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	for _, ledger := range ledgerInfo {
		var created *pojo.LedgerInfo
		if err := pgxscan.Get(
			context.Background(),
			tx,
			created,
			insertLedgerQuery,
			ledger.Owner,
			ledger.Name,
			ledger.TypeCode,
			ledger.Currency,
			ledger.Balance,
		); err != nil {
			return nil, err
		}
		createdLedgers = append(createdLedgers, created)
	}

	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}

	return createdLedgers, nil
}

func (l *LedgerDao) Update(userId string, ledger *pojo.UpdateLedgerInfo) (*pojo.UpdateLedgerInfo, error) {
	var updated *pojo.UpdateLedgerInfo
	updateLedgerQuery := `
	UPDATE user_ledgers 
	SET ledger_name = $2, ledger_status = $3, ledger_tags = $4
	WHERE ledger_id = $1
	RETURNING (
		ledger_id,
		ledger_name,
		ledger_status,
		ledger_tags
	)
	`
	if err := pgxscan.Get(
		context.Background(),
		l.pool,
		updated,
		updateLedgerQuery,
		ledger.Id, ledger.Name, ledger.LedgerStatus, ledger.LedgerTags,
	); err != nil {
		return updated, err
	}

	return updated, nil
}
