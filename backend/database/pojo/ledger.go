package pojo

import "time"

// TODO: Support ledger rules
type LedgerInfo struct {
	Id           string    `json:"ledgerId" db:"ledger_id"`
	Owner        string    `json:"ledgerOwner" db:"ledger_owner"`
	Name         string    `json:"ledgerName" db:"ledger_name"`
	TypeCode     string    `json:"ledgerTypeCode" db:"ledger_type_id"`
	TypeName     string    `json:"ledgerTypeName" db:"ledger_type_name"`
	Currency     string    `json:"currency" db:"currency"`
	CurrencyName string    `json:"currencyName" db:"currency_name"`
	Balance      string    `json:"balance" db:"balance"`
	LedgerTags   []string  `json:"ledgerTags" db:"ledger_tags"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
}

// TODO: Support ledger rules
type UpdateLedgerInfo struct {
	Name       string   `json:"ledgerName" db:"ledger_name"`
	TypeCode   string   `json:"ledgerTypeCode" db:"ledger_type_id"`
	LedgerTags []string `json:"ledgerTags" db:"ledger_tags"`
}
