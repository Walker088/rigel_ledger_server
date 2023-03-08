package protected

import (
	"encoding/json"
	"net/http"

	"github.com/Walker088/rigel_ledger_server/backend/database/dao"
	"github.com/Walker088/rigel_ledger_server/backend/database/pojo"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LedgerHandler struct {
	Dao *dao.LedgerDao
}

func NewLedgerHandler(pool *pgxpool.Pool) *LedgerHandler {
	return &LedgerHandler{
		Dao: dao.NewLedgerDao(pool),
	}
}

func (l *LedgerHandler) GetLedgerHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	ledgerId := chi.URLParam(r, "ledgerId")
	if ledger, err := l.Dao.Get(userId, ledgerId); err == nil {
		response, _ := json.Marshal(ledger)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func (l *LedgerHandler) GetLedgerLstHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	if ledgers, err := l.Dao.GetList(userId); err == nil {
		response, _ := json.Marshal(ledgers)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func (l *LedgerHandler) CreateLedgerHandler(w http.ResponseWriter, r *http.Request) {

	var ledgerLst []*pojo.LedgerInfo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ledgerLst); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if inserted, err := l.Dao.Create(ledgerLst); err == nil {
		response, _ := json.Marshal(inserted)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func (l *LedgerHandler) UpdateLedgerHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	var toUpdate pojo.UpdateLedgerInfo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&toUpdate); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if updated, err := l.Dao.Update(userId, &toUpdate); err == nil {
		response, _ := json.Marshal(updated)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}
