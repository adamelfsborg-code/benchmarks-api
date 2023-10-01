package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/adamelfsborg-code/benchmarks-api/internal/api/benchmark"
)

type App struct {
	router   http.Handler
	database benchmark.Actions
	config   Config
}

func Build(ctx context.Context) (*App, error) {
	c, err := loadConfig()
	if err != nil {
		fmt.Printf("Failed to build config: %v", err)
		return nil, err
	}

	var database benchmark.Actions

	dc := &DatabaseConn{
		Addr:     c.DBAddr,
		Database: c.DBDatabase,
		User:     c.DBUser,
		Password: c.DBPassword,
	}

	con, err := dc.loadDatabase(ctx)
	if err != nil {
		fmt.Printf("Failed to build database: %v", err)
		return nil, err
	}

	if c.Mode == "dev" {
		con.AddQueryHook(QueryLogger{})
	}

	database = &benchmark.Database{
		Client: con,
	}

	app := &App{
		config:   c,
		database: database,
	}

	app.loadRoutes()

	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.ServerPort),
		Handler: a.router,
	}

	fmt.Println("Starting Server!")

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("Failed to start server: %w", err)
		}

		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
