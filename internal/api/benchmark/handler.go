package benchmark

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Actions interface {
	Create(ctx context.Context, benchmark Benchmark) error
	List(ctx context.Context, page PageCursor) (FindResult, error)
}

type BenchmarkHandler struct {
	Actions Actions
}

type PageCursor struct {
	Size   uint64
	Offset uint64
}

type FindResult struct {
	Benchmarks []Benchmark
	Cursor     uint64
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
		Operation: body.Operation,
		Data:      body.Data,
	}

	err = b.Actions.Create(r.Context(), benchmark)
	if err != nil {
		fmt.Println("Failed to Insert: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (b *BenchmarkHandler) List(w http.ResponseWriter, r *http.Request) {
	cursorStr := r.URL.Query().Get("cursor")
	if cursorStr == "" {
		cursorStr = "0"
	}

	const decimal = 10
	const bitSize = 64
	cursor, err := strconv.ParseUint(cursorStr, decimal, bitSize)
	if err != nil {
		fmt.Println("Failed to parse cursor: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	const size = 50
	res, err := b.Actions.List(r.Context(), PageCursor{
		Offset: cursor,
		Size:   size,
	})

	if err != nil {
		fmt.Println("Failed to find all: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response struct {
		Benchmarks []Benchmark `json:"benchmarks"`
		Next       uint64      `json:"next,omitempty"`
	}

	response.Benchmarks = res.Benchmarks
	response.Next = res.Cursor

	data, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Failed to marshal: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
