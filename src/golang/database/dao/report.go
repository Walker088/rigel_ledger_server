package dao

import (
	"context"

	"github.com/Walker088/rigel_ledger_server/src/golang/database/pojo"
	"github.com/georgysavva/scany/v2/pgxscan"
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

func (r *ReportDao) GetAnnualBalanceSheet(userId string, year int) (*pojo.BalanceSheet, error) {
	var balanceSheet *pojo.BalanceSheet
	query := ``
	if err := pgxscan.Get(
		context.Background(),
		r.pool,
		&balanceSheet,
		query,
		userId,
		year,
	); err != nil {
		return balanceSheet, nil
	}
	return balanceSheet, nil
}

func (r *ReportDao) GetSeasonalBalanceSheet(userId string, year int, season int) (*pojo.BalanceSheet, error) {
	var balanceSheet *pojo.BalanceSheet
	query := ``
	if err := pgxscan.Get(
		context.Background(),
		r.pool,
		&balanceSheet,
		query,
		userId,
		year,
		season,
	); err != nil {
		return balanceSheet, nil
	}
	return balanceSheet, nil
}

func (r *ReportDao) GetAnnualIncomeStatement(userId string, year int) (*pojo.IncomeStatement, error) {
	var incomeStatement *pojo.IncomeStatement
	query := ``
	if err := pgxscan.Get(
		context.Background(),
		r.pool,
		&incomeStatement,
		query,
		userId,
		year,
	); err != nil {
		return incomeStatement, nil
	}
	return incomeStatement, nil
}

func (r *ReportDao) GetSeasonalIncomeStatement(userId string, year int, season int) (*pojo.IncomeStatement, error) {
	var incomeStatement *pojo.IncomeStatement
	query := ``
	if err := pgxscan.Get(
		context.Background(),
		r.pool,
		&incomeStatement,
		query,
		userId,
		year,
		season,
	); err != nil {
		return incomeStatement, nil
	}
	return incomeStatement, nil
}
