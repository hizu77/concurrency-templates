package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/hizu77/concurrency-templates/internal/infra/generator"
	"github.com/hizu77/concurrency-templates/internal/infra/workerpool"
	model "github.com/hizu77/concurrency-templates/internal/model/order"
	"github.com/hizu77/concurrency-templates/internal/service/order"
)

const (
	orderCount = 20
	workers    = 5
	fanLimit   = 5
)

func main() {
	// it might be better to write a local benchmark, but why :)
	// here we will use workerpool based on pipeline with fan on order processing
	start := time.Now()

	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	service := order.New(logger)

	output := workerpool.StartWithPipeline(
		ctx,
		generator.Generate(orderCount),
		func(a <-chan model.Order) <-chan model.Order {
			return service.FanPipeline(ctx, a, fanLimit)
		},
		workers,
	)

	completed := make([]model.Order, 0, orderCount)

	for i := range output {
		completed = append(completed, i)
	}

	fmt.Println(time.Since(start).Seconds()) // 7 is better than 13 :)
}
