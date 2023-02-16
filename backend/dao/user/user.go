package dao

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserDao struct {
	pool *pgxpool.Pool
}

type UserInfo struct {
	Id            string        `json:"userId"`
	Name          string        `json:"userName"`
	Mail          string        `json:"userMail"`
	TypeCode      int8          `json:"userTypeCode"`
	TypeName      string        `json:"userTypeName"`
	MainCountry   string        `json:"mainCountry"`
	MainCurrency  string        `json:"mainCurrency"`
	MainLanguage  string        `json:"mainLanguage"`
	LedgerSummary LedgerSummary `json:"ledgerSummary"`
}
type LedgerSummary struct {
	Balance struct {
		Assets      int `json:"assets"`
		Liabilities int `json:"liabilities"`
		NetWorth    int `json:"networth"`
	} `json:"balance"`
	Income struct {
		MonthlyIncome  int `json:"monthlyIncome"`
		MonthlyExpence int `json:"monthlyExpence"`
	} `json:"income"`
	Transactions struct {
		AvailableRange struct {
			FirstAt time.Time `json:"firstAt"`
			LastAt  time.Time `json:"lastAt"`
		} `json:"availableRange"`
		MonthlyTxs []MonthlyTxs `json:"monthlyTxs"`
	} `json:"transactions"`
}
type MonthlyTxs struct {
	TransacId         int       `json:"transacId"`
	TransacDate       time.Time `json:"transacDate"`
	DebitAccountCode  string    `json:"debitAccountCode"`
	DebitAccountName  string    `json:"debitAccountName"`
	CreditAccountCode string    `json:"creditAccountCode"`
	CreditAccountName string    `json:"creditAccountName"`
	Currency          string    `json:"currency"`
	OriginalAmt       int       `json:"originalAmt"`
	AdjustedAmt       int       `json:"adjustedAmt"`
}

func (d *UserDao) Get(userId string) (UserInfo, error) {
	var userInfo UserInfo
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
		return userInfo, err
	}

	monthlyTxs := make([]MonthlyTxs, 0)
	queryUserTx := `
	
	`
	rows, err := d.pool.Query(context.Background(), queryUserTx, userId)
	if err != nil {
		return userInfo, err
	}
	defer rows.Close()

	for rows.Next() {
		var tx struct {
			TransacId         int       `json:"transacId"`
			TransacDate       time.Time `json:"transacDate"`
			DebitAccountCode  string    `json:"debitAccountCode"`
			DebitAccountName  string    `json:"debitAccountName"`
			CreditAccountCode string    `json:"creditAccountCode"`
			CreditAccountName string    `json:"creditAccountName"`
			Currency          string    `json:"currency"`
			OriginalAmt       int       `json:"originalAmt"`
			AdjustedAmt       int       `json:"adjustedAmt"`
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
			return userInfo, err
		}
		monthlyTxs = append(monthlyTxs, tx)
	}
	userInfo.LedgerSummary.Transactions.MonthlyTxs = monthlyTxs
	return userInfo, nil
}
