package protected

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LedgerJournalDao struct {
	pool *pgxpool.Pool
}

func NewLedgerJournalHandler(p *pgxpool.Pool) *LedgerJournalDao {
	return &LedgerJournalDao{
		pool: p,
	}
}

func (l *LedgerJournalDao) GetJournalHandler(w http.ResponseWriter, r *http.Request) {

}

func (l *LedgerJournalDao) GetJournalLstHandler(w http.ResponseWriter, r *http.Request) {

}

func (l *LedgerJournalDao) CreateJournalHandler(w http.ResponseWriter, r *http.Request) {

}

func (l *LedgerJournalDao) RevertJournalHandler(w http.ResponseWriter, r *http.Request) {

}
