package dao

import (
	"context"

	"github.com/Walker088/rigel_ledger_server/backend/database/pojo"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserDao struct {
	pool *pgxpool.Pool
}

func NewUserDao(p *pgxpool.Pool) *UserDao {
	return &UserDao{
		pool: p,
	}
}

func (d *UserDao) GetByGithubId(ghId string) (*pojo.UserInfo, error) {
	var userInfo pojo.UserInfo
	queryUser := `
	SELECT
		user_id,
		user_name,
		user_mail,
		user_type user_type_code,
		CASE user_type
			WHEN 0 THEN 'Personal'
			WHEN 1 THEN 'Company'
			WHEN 2 THEN 'Pending'
			ELSE 'Unknown'
		END user_type_name,
		main_country,
		(SELECT country_name FROM ref_countries_iso3166_1 WHERE alphabetic_code_2 = main_country) main_country_name,
		main_currency,
		(SELECT currency_name FROM ref_currencies_iso4217 WHERE alphabetic_code = main_currency) main_currency_name,
		main_language,
		(SELECT language_name FROM ref_languages_iso639_1 WHERE alphabetic_code = main_language) main_language_name
	FROM
		user_info
	WHERE
		user_id_gh  = $1
	`
	if err := pgxscan.Get(context.Background(), d.pool, &userInfo, queryUser, ghId); err != nil {
		return &userInfo, err
	}
	return &userInfo, nil
}

func (d *UserDao) GetBasic(userId string) (*pojo.UserInfo, error) {
	var userInfo pojo.UserInfo
	queryUser := `
	SELECT
		user_id,
		user_name,
		user_mail,
		user_type user_type_code,
		CASE user_type
			WHEN 0 THEN 'Personal'
			WHEN 1 THEN 'Company'
			WHEN 2 THEN 'Pending'
			ELSE 'Unknown'
		END user_type_name,
		main_country,
		(SELECT country_name FROM ref_countries_iso3166_1 WHERE alphabetic_code_2 = main_country) main_country_name,
		main_currency,
		(SELECT currency_name FROM ref_currencies_iso4217 WHERE alphabetic_code = main_currency) main_currency_name,
		main_language,
		(SELECT language_name FROM ref_languages_iso639_1 WHERE alphabetic_code = main_language) main_language_name
	FROM
		user_info
	WHERE
		user_id_gh  = $1
	`
	if err := pgxscan.Get(context.Background(), d.pool, &userInfo, queryUser, userId); err != nil {
		return &userInfo, err
	}
	return &userInfo, nil
}

// TODO: Make the query being able to retrive amount_adjusted other than USD for users
func (d *UserDao) GetComplete(userId string) (*pojo.UserInfo, error) {
	userInfo, err := d.GetBasic(userId)
	if err != nil {
		return userInfo, err
	}

	var ledgerSummary pojo.LedgerSummary
	queryLedgerSummary := `
	SELECT
		x.bl_assets,
		x.bl_liabilities,
		x.bl_assets - x.bl_liabilities bl_networth,
		COALESCE(x.ie_monthly_income, 0) ie_monthly_income,
		COALESCE(x.ie_monthly_expense, 0) ie_monthly_expense,
		x.tx_first_at,
		x.tx_last_at
	FROM
	(
		SELECT
		(
			SELECT
				SUM(balance_in_usd)::NUMERIC(20, 6) balance
			FROM (
				SELECT
					l.currency,
					CASE l.currency
						WHEN 'USD' THEN SUM(l.balance)
						ELSE SUM(l.balance) / (SELECT exchange_rate FROM ref_exchange_rates WHERE from_code = 'USD' AND to_code = l.currency)
					END balance_in_usd
				FROM 
					user_ledgers l 
				WHERE 
					l.ledger_owner = $1
					AND l.ledger_type_id LIKE '1%'
					AND l.ledger_type_id NOT IN (
						'1139', '1229', '1419', '1438', '1439',
						'1448', '1449', '1518', '1519', '1569', 
						'1588', '1589', '1719', '1729', '1749', 
						'1758', '1759', '1769', '1859')
				GROUP BY l.currency
			) x
		) bl_assets,
		(
			SELECT
				SUM(balance_in_usd)::NUMERIC(20, 6) balance
			FROM (
				SELECT
					CASE l.currency
						WHEN 'USD' THEN SUM(l.balance)
						ELSE SUM(l.balance) / (SELECT exchange_rate FROM ref_exchange_rates WHERE from_code = 'USD' AND to_code = l.currency)
					END balance_in_usd
				FROM 
					user_ledgers l 
				WHERE 
					l.ledger_owner = $1
					AND l.ledger_type_id LIKE '2%'
				GROUP BY l.currency
			) x
		) bl_liabilities,
		(
			SELECT
				SUM(balance_in_usd)::NUMERIC(20, 6) balance
			FROM (
				SELECT
					CASE l.currency
						WHEN 'USD' THEN SUM(j.amount)
						ELSE SUM(j.amount) / (SELECT exchange_rate FROM ref_exchange_rates WHERE from_code = 'USD' AND to_code = l.currency)
					END balance_in_usd
				FROM 
					user_ledger_transactions t
				LEFT JOIN user_ledgers l USING (ledger_id)
				LEFT JOIN user_ledger_journal j ON t.ledger_id = j.debit_account 
				WHERE 
					l.ledger_owner = $1
					AND l.ledger_type_id LIKE 'A%'
					AND EXTRACT (MONTH FROM t.created_at) = EXTRACT (MONTH FROM CURRENT_DATE)
				GROUP BY l.currency
			) x	
		) ie_monthly_income,
		(
			SELECT
				SUM(balance_in_usd)::NUMERIC(20, 6) balance
			FROM (
				SELECT
					CASE l.currency
						WHEN 'USD' THEN SUM(j.amount)
						ELSE SUM(j.amount) / (SELECT exchange_rate FROM ref_exchange_rates WHERE from_code = 'USD' AND to_code = l.currency)
					END balance_in_usd
				FROM 
					user_ledger_transactions t
				LEFT JOIN user_ledgers l USING (ledger_id)
				LEFT JOIN user_ledger_journal j ON t.ledger_id = j.credit_account  
				WHERE 
					l.ledger_owner = $1
					AND l.ledger_type_id LIKE 'B%'
					AND EXTRACT (MONTH FROM t.created_at) = EXTRACT (MONTH FROM CURRENT_DATE)
				GROUP BY l.currency
			) x	
		) ie_monthly_expense,
		MIN(j.transac_date) tx_first_at,
		MAX(j.transac_date) tx_last_at
		FROM 
			user_ledger_journal j
		WHERE 
			j.user_id = $1
			AND EXTRACT (MONTH FROM j.transac_date) = EXTRACT (MONTH FROM CURRENT_DATE)
	) x

	`
	if err := pgxscan.Get(context.Background(), d.pool, &ledgerSummary, queryLedgerSummary, userId); err != nil {
		return userInfo, err
	}
	userInfo.LedgerSummary = ledgerSummary

	monthlyTxs := make([]pojo.MonthlyTxs, 0)
	queryUserTx := `
	SELECT
		transac_id,
		transac_date,
		credit_account,
		credit_account_name,
		debit_account,
		debit_account_name,
		currency,
		amount_origin,
		CASE currency
			WHEN 'USD' THEN amount_origin
			ELSE amount_origin / (SELECT exchange_rate FROM ref_exchange_rates WHERE from_code = 'USD' AND to_code = currency)
		END amount_adjusted
	FROM (
		SELECT
			j.transac_id,
			j.transac_date,
			j.credit_account,
			(SELECT ledger_name FROM user_ledgers WHERE ledger_id = j.credit_account) credit_account_name,
			j.debit_account,
			(SELECT ledger_name FROM user_ledgers WHERE ledger_id = j.debit_account) debit_account_name,
			(SELECT currency FROM user_ledgers WHERE ledger_id IN (j.credit_account, j.debit_account) GROUP BY currency) currency,
			j.amount amount_origin
		FROM
			user_ledger_journal j 
		WHERE 
			user_id = $1
			AND EXTRACT (MONTH FROM j.transac_date) = EXTRACT (MONTH FROM CURRENT_DATE)
	) x
	`
	if err := pgxscan.Select(context.Background(), d.pool, &monthlyTxs, queryUserTx, userId); err != nil {
		return userInfo, err
	}
	userInfo.LedgerSummary.Transactions.MonthlyTxs = monthlyTxs

	return userInfo, nil
}

func (d *UserDao) Update(userId string, u *pojo.UpdateUserInfo) (*pojo.UpdateUserInfo, error) {
	updateUserQuery := `
	UPDATE user_info
	SET 
		user_name  = $1,
		user_mail = $2,
		main_country = $3,
		main_currency = $4,
		main_language = $5
	WHERE
		user_id = $6
	RETURNING (user_id, user_name, user_mail, main_country, main_currency, main_language)
	`

	if err := pgxscan.Get(
		context.Background(),
		d.pool,
		u,
		updateUserQuery,
		u.Name, u.Mail, u.MainCountry,
		u.MainCurrency, u.MainLanguage, userId,
	); err != nil {
		return u, err
	}

	return u, nil
}
