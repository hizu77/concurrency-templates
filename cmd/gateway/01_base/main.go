package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/hizu77/concurrency-templates/internal/infra/generator"
	model "github.com/hizu77/concurrency-templates/internal/model/order"
	"github.com/hizu77/concurrency-templates/internal/service/order"
)

const orderCount = 20

func main() {
	// it might be better to write a local benchmark, but why :)
	start := time.Now()

	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	service := order.New(logger)
	generated := generator.Generate(orderCount)

	started := make([]model.Order, 0, len(generated))
	for order := range generated {
		started = append(started, service.Start(ctx, order))
	}

	processed := make([]model.Order, 0, len(started))
	for _, order := range started {
		processed = append(processed, service.Process(ctx, order))
	}

	completed := make([]model.Order, 0, len(processed))
	for _, order := range processed {
		completed = append(completed, service.Complete(ctx, order))
	}

	fmt.Println(time.Since(start).Seconds()) // 80 so bad
}
