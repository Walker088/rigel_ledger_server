package dao

import (
	"context"
	"time"

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
	err := pgxscan.Get(context.Background(), d.pool, &userInfo, queryUser, ghId)
	if err != nil {
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
	err := pgxscan.Get(context.Background(), d.pool, &userInfo, queryUser, userId)
	if err != nil {
		return &userInfo, err
	}
	return &userInfo, nil
}

func (d *UserDao) GetComplete(userId string) (*pojo.UserInfo, error) {
	var userInfo pojo.UserInfo
	queryUser := `
	
	`
	err := d.pool.QueryRow(context.Background(), queryUser, userId).Scan(
		&userInfo.Id,
		&userInfo.Name,
		&userInfo.Mail,
		&userInfo.TypeCode,
		&userInfo.TypeName,
		&userInfo.MainCountry,
		&userInfo.MainCurrency,
		&userInfo.MainLanguage,
		&userInfo.LedgerSummary.Balance.Assets,
		&userInfo.LedgerSummary.Balance.Liabilities,
		&userInfo.LedgerSummary.Balance.NetWorth,
		&userInfo.LedgerSummary.Income.MonthlyIncome,
		&userInfo.LedgerSummary.Income.MonthlyExpence,
		&userInfo.LedgerSummary.Transactions.AvailableRange.FirstAt,
		&userInfo.LedgerSummary.Transactions.AvailableRange.LastAt,
	)
	if err != nil {
		return &userInfo, err
	}

	monthlyTxs := make([]pojo.MonthlyTxs, 0)
	queryUserTx := `
	
	`
	rows, err := d.pool.Query(context.Background(), queryUserTx, userId)
	if err != nil {
		return &userInfo, err
	}
	defer rows.Close()

	for rows.Next() {
		var tx struct {
			TransacId         int       `json:"transacId" db:"transac_id"`
			TransacDate       time.Time `json:"transacDate" db:"transac_date"`
			DebitAccountCode  string    `json:"debitAccountCode" db:"debit_account"`
			DebitAccountName  string    `json:"debitAccountName" db:"debit_account_name"`
			CreditAccountCode string    `json:"creditAccountCode" db:"credit_account"`
			CreditAccountName string    `json:"creditAccountName" db:"credit_account_name"`
			Currency          string    `json:"currency" db:"currency"`
			OriginalAmt       string    `json:"originalAmt" db:"amount_origin"`
			AdjustedAmt       string    `json:"adjustedAmt" db:"amount_adjusted"`
		}
		if err := rows.Scan(
			&tx.TransacId,
			&tx.TransacDate,
			&tx.DebitAccountCode,
			&tx.DebitAccountName,
			&tx.CreditAccountCode,
			&tx.CreditAccountName,
			&tx.Currency,
			&tx.OriginalAmt,
			&tx.AdjustedAmt,
		); err != nil {
			return &userInfo, err
		}
		monthlyTxs = append(monthlyTxs, tx)
	}
	userInfo.LedgerSummary.Transactions.MonthlyTxs = monthlyTxs
	return &userInfo, nil
}

func (d *UserDao) Update(userId string, u *pojo.UserInfo) (*pojo.UserInfo, error) {

	updateUserQuery := `
	
	`
	d.pool.Exec(context.Background(), updateUserQuery,
		u.Name, u.Mail, u.MainCountry, u.MainCurrency, u.MainLanguage,
	)
	return u, nil
}
