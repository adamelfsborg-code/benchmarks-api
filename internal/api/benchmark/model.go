package benchmark

type Benchmark struct {
	Head BenchmarkHead
	Data []BenchmarkRow
}

type BenchmarkHead struct {
	Operation string   `json:"operation"`
	Timestamp string   `json:"timestamp"`
	Tags      []string `json:"tags"`
}

type BenchmarkRow struct {
	Xlabel string `json:"x_label"`
	Ylabel string `json:"y_label"`

	Xvalue uint64 `json:"x_value"`
	Yvalue uint64 `json:"y_value"`

	Xunit string `json:"x_unit"`
	Yunit string `json:"y_unit"`
}
