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

const (
	orderCount = 20
	fanLimit   = 5
)

func main() {
	// it might be better to write a local benchmark, but why :)
	// here we will use pipeline in order processing and fan on process stage
	start := time.Now()

	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	service := order.New(logger)

	pipeline := service.FanPipeline(ctx, generator.Generate(orderCount), fanLimit)
	completed := make([]model.Order, 0, orderCount)

	for i := range pipeline {
		completed = append(completed, i)
	}

	fmt.Println(time.Since(start).Seconds()) // 23 is better than 61 :)
}
