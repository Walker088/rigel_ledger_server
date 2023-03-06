package dao

import (
	"context"
	"fmt"

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
	SELECT
		j.user_id,
		j.transac_id,
		j.transac_date,
		j.credit_account,
		cd.ledger_name credit_account_name,
		j.debit_account,
		db.ledger_name debit_account_name,
		cd.currency,
		j.amount amount_origin,
		CASE
			WHEN u.main_currency = cd.currency THEN j.amount
			ELSE j.amount 
				/ (SELECT exchange_rate FROM ref_exchange_rates WHERE from_code = 'USD' AND to_code = cd.currency)
				* (SELECT exchange_rate FROM ref_exchange_rates WHERE from_code = 'USD' AND to_code = u.main_currency)
		END amount_adjusted,
		j.photo_addr,
		j.description,
		j.created_at 
	FROM 
		user_ledger_journal j
	JOIN user_ledgers cd ON (cd.ledger_id = j.credit_account)
	JOIN user_ledgers db ON (db.ledger_id = j.credit_account)
	JOIN user_info u ON (u.user_id = j.user_id)
	WHERE 
		j.user_id = $1 AND j.transac_id = $2
	`

	if err := pgxscan.Get(
		context.Background(),
		l.pool,
		&ledger,
		queryLedgerInfo,
		userId, transacId,
	); err != nil {
		return ledger, err
	}
	return ledger, nil
}

func (l *LedgerJournalDao) GetList(userId string, pageSize int, page int) ([]*pojo.LedgerJournal, error) {
	var ledger []*pojo.LedgerJournal
	limit := pageSize
	offset := (page - 1) * pageSize

	queryLedgerInfo := `
	SELECT
		j.user_id,
		j.transac_id,
		j.transac_date,
		j.credit_account,
		cd.ledger_name credit_account_name,
		j.debit_account,
		db.ledger_name debit_account_name,
		cd.currency,
		j.amount amount_origin,
		CASE
			WHEN u.main_currency = cd.currency THEN j.amount
			ELSE j.amount 
				/ (SELECT exchange_rate FROM ref_exchange_rates WHERE from_code = 'USD' AND to_code = cd.currency)
				* (SELECT exchange_rate FROM ref_exchange_rates WHERE from_code = 'USD' AND to_code = u.main_currency)
		END amount_adjusted,
		j.photo_addr,
		j.description,
		j.created_at 
	FROM 
		user_ledger_journal j
	JOIN user_ledgers cd ON (cd.ledger_id = j.credit_account)
	JOIN user_ledgers db ON (db.ledger_id = j.credit_account)
	JOIN user_info u ON (u.user_id = j.user_id)
	WHERE 
		j.user_id = $1
	ORDER BY 
		j.transac_id DESC LIMIT $2 OFFSET $3
	`
	if err := pgxscan.Select(
		context.Background(),
		l.pool,
		&ledger,
		queryLedgerInfo,
		userId, limit, offset,
	); err != nil {
		return nil, err
	}

	return ledger, nil
}

func (l *LedgerJournalDao) Create(toCreateLst []*pojo.CreateJournalEntry) ([]*pojo.LedgerJournal, error) {
	var createdEntries = make([]*pojo.LedgerJournal, 0)
	insertJournalQuery := `
	INSERT INTO user_ledger_journal (user_id, transac_id, transac_date, credit_account, debit_account, amount, photo_addr, description)
	SELECT
		$1,
		COUNT(1) + 1,
		CURRENT_DATE,
		$2,
		$3,
		$4,
		$5,
		$6
	FROM
		user_ledger_journal
	WHERE user_id = $1
	RETURNING (user_id, transac_date, credit_account, debit_account, amount, photo_addr, description);
	`
	insertCreditAccTx := `
	INSERT INTO user_ledger_transactions (ledger_id, transac_id, balance)
	SELECT
		ledger_id,
		$2,
		CASE
			WHEN t.first_grade IN ('1', '5', '6', 'B') THEN balance + $3
			WHEN t.first_grade IN ('2', '3', '4', 'A') THEN balance - $3
			WHEN t.first_grade = '7' AND t.second_grade IN ('75', '76', '77', '78') THEN balance + $3
			WHEN t.first_grade = '7' AND t.second_grade IN ('71', '72', '73', '74') THEN balance - $3
		END
	FROM
		user_ledgers
	JOIN ref_ledger_types t USING (ledger_type_id)
	WHERE
		ledger_id = $1
	`
	insertDebitAccTx := `
	INSERT INTO user_ledger_transactions (ledger_id, transac_id, balance)
	SELECT
		ledger_id,
		$2,
		CASE
			WHEN t.first_grade IN ('1', '5', '6', 'B') THEN balance - $3
			WHEN t.first_grade IN ('2', '3', '4', 'A') THEN balance + $3
			WHEN t.first_grade = '7' AND t.second_grade IN ('75', '76', '77', '78') THEN balance - $3
			WHEN t.first_grade = '7' AND t.second_grade IN ('71', '72', '73', '74') THEN balance + $3
		END
	FROM
		user_ledgers
	JOIN ref_ledger_types t USING (ledger_type_id)
	WHERE
		ledger_id = $1
	`

	tx, err := l.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	for _, toCreate := range toCreateLst {
		var created *pojo.LedgerJournal
		if err := pgxscan.Get(
			context.Background(),
			tx,
			created,
			insertJournalQuery,
			toCreate.UserId, toCreate.CreditAccountCode, toCreate.DebitAccountCode,
			toCreate.Amount, toCreate.Photos, toCreate.Description,
		); err != nil {
			return nil, err
		}

		if _, err := tx.Exec(
			context.Background(),
			insertCreditAccTx,
			toCreate.CreditAccountCode, created.TransacId, toCreate.Amount,
		); err != nil {
			return nil, err
		}

		if _, err := tx.Exec(
			context.Background(),
			insertDebitAccTx,
			toCreate.DebitAccountCode, created.TransacId, toCreate.Amount,
		); err != nil {
			return nil, err
		}
		createdEntries = append(createdEntries, created)
	}

	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}

	return createdEntries, nil
}

func (l *LedgerJournalDao) RevertTx(userId string, transacId int) ([]*pojo.LedgerJournal, error) {
	origin, err := l.Get(userId, transacId)
	if err != nil {
		return nil, nil
	}
	revertDesc := fmt.Sprintf("Reverted origin transaction %s, affecting %s and %s", origin.TransacId, origin.CreditAccountName, origin.DebitAccountName)
	toRevert := &pojo.CreateJournalEntry{
		UserId:            userId,
		CreditAccountCode: origin.DebitAccountCode,
		DebitAccountCode:  origin.CreditAccountCode,
		Currency:          origin.Currency,
		Amount:            origin.OriginalAmt,
		Description:       revertDesc,
		Photos:            origin.PhotoAddr,
	}
	reverted, err := l.Create([]*pojo.CreateJournalEntry{toRevert})
	if err != nil {
		return nil, nil
	}
	return reverted, nil
}
