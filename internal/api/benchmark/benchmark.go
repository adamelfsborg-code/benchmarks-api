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

	_, err := d.Client.Model(&benchmark.Head).Returning("id").Insert()
	if err != nil {
		return fmt.Errorf("Failed to create benchmark: %w", err)
	}

	fmt.Printf("Generated Benchmark Head ID: %s\n", benchmark.Head.ID)

	benchmarkRow := BenchmarkRow{
		BenchmarkID: benchmark.Head.ID,
		Data:        benchmark.Row.Data,
	}

	d.Client.Model(&benchmarkRow).Returning("id").Insert()

	fmt.Printf("Generated Benchmark Row ID: %s\n", benchmarkRow.ID)

	return nil
}
