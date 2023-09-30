package benchmark

import (
	"net/http"

	"github.com/go-pg/pg/v10"
)

type BenchmarkHandler struct {
	Repo *pg.DB
}

func (b *BenchmarkHandler) Create(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
