package benchmark

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Actions interface {
	Create(ctx context.Context, benchmark Benchmark) error
}

type BenchmarkHandler struct {
	Actions Actions
}

func (b *BenchmarkHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Operation string          `json:"operation"`
		Data      json.RawMessage `json:"data"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	benchmark := Benchmark{
		Head: BenchmarkHead{
			Operation: body.Operation,
		},
		Row: BenchmarkRow{
			Data: body.Data,
		},
	}

	err = b.Actions.Create(r.Context(), benchmark)
	if err != nil {
		fmt.Println("Failed to Insert: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
