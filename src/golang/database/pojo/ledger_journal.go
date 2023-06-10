package pojo

import "time"

type LedgerJournal struct {
	UserId            string    `json:"userId" db:"user_id"`
	TransacId         string    `json:"transacId" db:"transac_id"`
	TransacDate       time.Time `json:"transacDate" db:"transac_date"`
	CreditAccountCode string    `json:"creditAccountCode" db:"credit_account"`
	CreditAccountName string    `json:"creditAccountName" db:"credit_account_name"`
	DebitAccountCode  string    `json:"debitAccountCode" db:"debit_account"`
	DebitAccountName  string    `json:"debitAccountName" db:"debit_account_name"`
	Currency          string    `json:"currency" db:"currency"`
	OriginalAmt       string    `json:"originalAmt" db:"amount_origin"`
	AdjustedAmt       string    `json:"adjustedAmt" db:"amount_adjusted"`
	PhotoAddr         []string  `json:"photoAddr" db:"photo_addr"`
	Description       string    `json:"description" db:"description"`
	CreatedAt         time.Time `json:"createdAt" db:"created_at"`
}

type CreateJournalEntry struct {
	UserId            string    `json:"userId" db:"user_id"`
	TransacDate       time.Time `json:"transacDate" db:"transac_date"`
	CreditAccountCode string    `json:"creditAccountCode" db:"credit_account"`
	DebitAccountCode  string    `json:"debitAccountCode" db:"debit_account"`
	Currency          string    `json:"currency" db:"currency"`
	Amount            string    `json:"amount" db:"amount"`
	Description       string    `json:"description" db:"description"`
	Photos            []string  `json:"photos" db:"photo_addr"`
}
