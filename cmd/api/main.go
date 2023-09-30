package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/adamelfsborg-code/benchmarks-api/internal/api/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app, err := app.Build(ctx)
	if err != nil {
		log.Fatal("Failed to load API", err)
	}
	err = app.Start(ctx)
	if err != nil {
		log.Fatal("Failed to start API", err)
	}
}
