package pojo

import "time"

type UserInfo struct {
	Id               string        `json:"userId" db:"user_id"`
	Name             string        `json:"userName" db:"user_name"`
	Mail             string        `json:"userMail" db:"user_mail"`
	TypeCode         uint8         `json:"userTypeCode" db:"user_type_code"`
	TypeName         string        `json:"userTypeName" db:"user_type_name"`
	MainCountry      string        `json:"mainCountry" db:"main_country"`
	MainCountryName  string        `json:"mainCountryName" db:"main_country_name"`
	MainCurrency     string        `json:"mainCurrency" db:"main_currency"`
	MainCurrencyName string        `json:"mainCurrencyName" db:"main_currency_name"`
	MainLanguage     string        `json:"mainLanguage" db:"main_language"`
	MainLanguageName string        `json:"mainLanguageName" db:"main_language_name"`
	LedgerSummary    LedgerSummary `json:"ledgerSummary" db:""`
}
type LedgerSummary struct {
	Balance struct {
		Assets      string `json:"assets" db:"bl_assets"`
		Liabilities string `json:"liabilities" db:"bl_liabilities"`
		NetWorth    string `json:"networth" db:"bl_networth"`
	} `json:"balance" db:""`
	Income struct {
		MonthlyIncome  string `json:"monthlyIncome" db:"ie_monthly_income"`
		MonthlyExpence string `json:"monthlyExpence" db:"ie_monthly_expense"`
	} `json:"income" db:""`
	Transactions struct {
		AvailableRange struct {
			FirstAt time.Time `json:"firstAt" db:"tx_first_at"`
			LastAt  time.Time `json:"lastAt" db:"tx_last_at"`
		} `json:"availableRange" db:""`
		MonthlyTxs []MonthlyTxs `json:"monthlyTxs"`
	} `json:"transactions" db:""`
}
type MonthlyTxs struct {
	TransacId         int       `json:"transacId" db:"transac_id"`
	TransacDate       time.Time `json:"transacDate" db:"transac_date"`
	CreditAccountCode string    `json:"creditAccountCode" db:"credit_account"`
	CreditAccountName string    `json:"creditAccountName" db:"credit_account_name"`
	DebitAccountCode  string    `json:"debitAccountCode" db:"debit_account"`
	DebitAccountName  string    `json:"debitAccountName" db:"debit_account_name"`
	Currency          string    `json:"currency" db:"currency"`
	OriginalAmt       string    `json:"originalAmt" db:"amount_origin"`
	AdjustedAmt       string    `json:"adjustedAmt" db:"amount_adjusted"`
}

type UpdateUserInfo struct {
	Id           string `json:"userId" db:"user_id"`
	Name         string `json:"userName" db:"user_name"`
	Mail         string `json:"userMail" db:"user_mail"`
	TypeCode     uint8  `json:"userTypeCode" db:"user_type_code"`
	MainCountry  string `json:"mainCountry" db:"main_country"`
	MainCurrency string `json:"mainCurrency" db:"main_currency"`
	MainLanguage string `json:"mainLanguage" db:"main_language"`
}
