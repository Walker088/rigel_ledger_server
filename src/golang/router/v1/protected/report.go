package protected

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Walker088/rigel_ledger_server/src/golang/database/dao"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ReportHandler struct {
	Dao *dao.ReportDao
}

func NewReportHandler(pool *pgxpool.Pool) *ReportHandler {
	return &ReportHandler{
		Dao: dao.NewReportDao(pool),
	}
}

func (rh *ReportHandler) GetAnnualBalanceSheet(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	year, err := strconv.Atoi(chi.URLParam(r, "year"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if year > time.Now().Year() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if report, err := rh.Dao.GetAnnualBalanceSheet(userId, year); err == nil {
		response, _ := json.Marshal(report)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func (rh *ReportHandler) GetSeasonalBalanceSheet(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	year, err := strconv.Atoi(chi.URLParam(r, "year"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	season, err := strconv.Atoi(chi.URLParam(r, "season"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if year > time.Now().Year() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if season < 1 || season > 4 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if report, err := rh.Dao.GetSeasonalBalanceSheet(userId, year, season); err == nil {
		response, _ := json.Marshal(report)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func (rh *ReportHandler) GetAnnualIncomeStatement(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	year, err := strconv.Atoi(chi.URLParam(r, "year"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if year > time.Now().Year() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if report, err := rh.Dao.GetAnnualIncomeStatement(userId, year); err == nil {
		response, _ := json.Marshal(report)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func (rh *ReportHandler) GetSeasonalIncomeStatement(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	year, err := strconv.Atoi(chi.URLParam(r, "year"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	season, err := strconv.Atoi(chi.URLParam(r, "season"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if year > time.Now().Year() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if season < 1 || season > 4 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if report, err := rh.Dao.GetSeasonalIncomeStatement(userId, year, season); err == nil {
		response, _ := json.Marshal(report)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

	w.WriteHeader(http.StatusInternalServerError)
}
