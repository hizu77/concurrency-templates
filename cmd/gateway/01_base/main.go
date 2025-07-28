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
		o, err := service.Start(ctx, order)
		if err != nil {
			logger.Error("service.Start", slog.Any("error", err))
			os.Exit(-1)
		}

		started = append(started, o)
	}

	processed := make([]model.Order, 0, len(started))
	for _, order := range started {
		o, err := service.Process(ctx, order)
		if err != nil {
			logger.Error("service.Start", slog.Any("error", err))
			os.Exit(-1)
		}

		processed = append(processed, o)
	}

	completed := make([]model.Order, 0, len(processed))
	for _, order := range processed {
		o, err := service.Complete(ctx, order)
		if err != nil {
			logger.Error("service.Start", slog.Any("error", err))
			os.Exit(-1)
		}

		completed = append(completed, o)
	}

	fmt.Println(time.Since(start).Seconds()) // 80 so bad
}
