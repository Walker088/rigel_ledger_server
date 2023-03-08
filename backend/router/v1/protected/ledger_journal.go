package protected

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Walker088/rigel_ledger_server/backend/database/dao"
	"github.com/Walker088/rigel_ledger_server/backend/database/pojo"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LedgerJournalHandler struct {
	Dao          *dao.LedgerJournalDao
	MIN_PageSize int
	MAX_PageSize int
}

func NewLedgerJournalHandler(pool *pgxpool.Pool) *LedgerJournalHandler {
	return &LedgerJournalHandler{
		Dao:          dao.NewLedgerJournalDao(pool),
		MIN_PageSize: 10,
		MAX_PageSize: 100,
	}
}

func (l *LedgerJournalHandler) GetJournalHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	transacId, err := strconv.Atoi(chi.URLParam(r, "transacId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if ledger, err := l.Dao.Get(userId, transacId); err == nil {
		response, _ := json.Marshal(ledger)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func (l *LedgerJournalHandler) GetJournalLstHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pagesize"))
	if pageSize < l.MIN_PageSize {
		pageSize = l.MIN_PageSize
	}
	if pageSize > l.MAX_PageSize {
		pageSize = l.MAX_PageSize
	}

	if journals, err := l.Dao.GetList(userId, pageSize, page); err == nil {
		response, _ := json.Marshal(journals)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func (l *LedgerJournalHandler) CreateJournalHandler(w http.ResponseWriter, r *http.Request) {
	var txLst []*pojo.CreateJournalEntry
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&txLst); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if inserted, err := l.Dao.Create(txLst); err == nil {
		response, _ := json.Marshal(inserted)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func (l *LedgerJournalHandler) RevertJournalHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	transacId, err := strconv.Atoi(chi.URLParam(r, "transacId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if updated, err := l.Dao.RevertTx(userId, transacId); err == nil {
		response, _ := json.Marshal(updated)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}
