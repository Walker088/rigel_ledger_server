package protected

import (
	"net/http"

	"github.com/Walker088/rigel_ledger_server/backend/database/dao"
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

}

func (l *LedgerHandler) GetLedgerLstHandler(w http.ResponseWriter, r *http.Request) {

}

func (l *LedgerHandler) CreateLedgerHandler(w http.ResponseWriter, r *http.Request) {

}

func (l *LedgerHandler) UpdateLedgerHandler(w http.ResponseWriter, r *http.Request) {

}
