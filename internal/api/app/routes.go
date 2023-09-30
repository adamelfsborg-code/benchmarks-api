package app

import (
	"net/http"

	"github.com/adamelfsborg-code/benchmarks-api/internal/api/benchmark"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/benchmarks", a.LoadBenchmarkRoutes)

	a.router = router

}

func (a *App) LoadBenchmarkRoutes(router chi.Router) {
	benchmarkHandler := &benchmark.BenchmarkHandler{
		Repo: a.database.Client,
	}
	router.Post("/", benchmarkHandler.Create)
}