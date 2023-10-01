package benchmark

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type Database struct {
	Client *pg.DB
}

func (d *Database) Create(ctx context.Context, benchmark Benchmark) error {
	benchmarkHead := BenchmarkHead{
		Operation: benchmark.Operation,
	}
	_, err := d.Client.Model(&benchmarkHead).Returning("id").Insert()
	if err != nil {
		return fmt.Errorf("Failed to create benchmark: %w", err)
	}

	benchmarkRow := BenchmarkRow{
		BenchmarkID: benchmarkHead.ID,
		Data:        benchmark.Data,
	}

	d.Client.Model(&benchmarkRow).Returning("id").Insert()

	return nil
}

func (d *Database) List(ctx context.Context, page PageCursor) (FindResult, error) {
	var benchmarks []Benchmark

	sqlQuery := `
		SELECT bmh.id, bmh.operation, bmh.timestamp, bmr.data
		FROM benchmark_heads bmh
		JOIN benchmark_rows bmr ON bmh.id = bmr.benchmark_id
		LIMIT 50
	`

	_, err := d.Client.Query(&benchmarks, sqlQuery)

	if err != nil {
		return FindResult{}, err
	}

	return FindResult{
		Benchmarks: benchmarks,
		Cursor:     page.Offset + page.Size,
	}, nil
}
