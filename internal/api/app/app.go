package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/adamelfsborg-code/test-api/internal/api/config"
	"github.com/adamelfsborg-code/test-api/internal/api/database"
	"github.com/adamelfsborg-code/test-api/internal/api/router"
)

type App struct {
	Router   router.Router
	Database database.Database
	Config   config.Config
}

func Build(ctx context.Context) (*App, error) {
	r, err := router.Build()
	if err != nil {
		fmt.Printf("Failed to build router: %v", err)
		return nil, err
	}

	c, err := config.Build()
	if err != nil {
		fmt.Printf("Failed to build config: %v", err)
		return nil, err
	}

	dc := &database.DatabaseConn{
		Addr:     c.DBAddr,
		Database: c.DBDatabase,
		User:     c.DBUser,
		Password: c.DBPassword,
	}

	d, err := dc.Build(ctx)
	if err != nil {
		fmt.Printf("Failed to build database: %v", err)
		return nil, err
	}

	app := &App{
		Router:   r,
		Config:   c,
		Database: *d,
	}

	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.Config.ServerPort),
		Handler: a.Router.Handler,
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
