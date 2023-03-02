package pojo

import "time"

// TODO: Support ledger rules
type LedgerInfo struct {
	Id              string    `json:"ledgerId" db:"ledger_id"`
	Owner           string    `json:"ledgerOwner" db:"ledger_owner"`
	OwnerName       string    `json:"ledgerOwnerName" db:"ledger_owner_name"`
	Name            string    `json:"ledgerName" db:"ledger_name"`
	TypeCode        string    `json:"ledgerTypeCode" db:"ledger_type_id"`
	TypeName        string    `json:"ledgerTypeName" db:"ledger_type_name"`
	Currency        string    `json:"currency" db:"currency"`
	CurrencyName    string    `json:"currencyName" db:"currency_name"`
	Balance         string    `json:"balance" db:"balance"`
	BalanceAdjusted string    `json:"balanceAdjusted" db:"balance_adjusted"`
	LedgerTags      []string  `json:"ledgerTags" db:"ledger_tags"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time `json:"updatedAt" db:"updated_at"`
}

// TODO: Support ledger rules
type UpdateLedgerInfo struct {
	Id           string   `json:"ledgerId" db:"ledger_id"`
	Name         string   `json:"ledgerName" db:"ledger_name"`
	LedgerStatus string   `json:"ledgerStatus" db:"ledger_status"`
	LedgerTags   []string `json:"ledgerTags" db:"ledger_tags"`
}
