package benchmark

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Benchmark struct {
	Head BenchmarkHead `json:"head"`
	Row  BenchmarkRow  `json:"row"`
}

type BenchmarkHead struct {
	ID        uuid.UUID `json:"id"`
	Operation string    `json:"operation"`
	Timestamp time.Time `json:"timestamp"`
}

type BenchmarkRow struct {
	ID          uuid.UUID       `json:"id"`
	BenchmarkID uuid.UUID       `json:"benchmark_id"`
	Data        json.RawMessage `json:"data"`
	Timestamp   time.Time       `json:"timestamp"`
}
