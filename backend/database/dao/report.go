package dao

import (
	"github.com/Walker088/rigel_ledger_server/backend/database/pojo"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ReportDao struct {
	pool *pgxpool.Pool
}

func NewReportDao(p *pgxpool.Pool) *ReportDao {
	return &ReportDao{
		pool: p,
	}
}

func (r *ReportDao) GetAnnualBalanceSheet(year int) (*pojo.BalanceSheet, error) {
	var balanceSheet *pojo.BalanceSheet
	return balanceSheet, nil
}

func (r *ReportDao) GetSeasonalBalanceSheet(year int, season int) (*pojo.BalanceSheet, error) {
	var balanceSheet *pojo.BalanceSheet
	return balanceSheet, nil
}

func (r *ReportDao) GetAnnualIncomeStatement(year int) (*pojo.IncomeStatement, error) {
	var incomeStatement *pojo.IncomeStatement
	return incomeStatement, nil
}

func (r *ReportDao) GetSeasonalIncomeStatement(year int, season int) (*pojo.IncomeStatement, error) {
	var incomeStatement *pojo.IncomeStatement
	return incomeStatement, nil
}
