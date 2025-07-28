package order

import (
	"context"

	"github.com/hizu77/concurrency-templates/internal/infra/fan"
	"github.com/hizu77/concurrency-templates/internal/infra/pipeline"
	model "github.com/hizu77/concurrency-templates/internal/model/order"
)

func (o *serviceImpl) FanPipeline(
	ctx context.Context,
	orders <-chan model.Order,
	fanLimit int,
) <-chan model.Order {
	createdToStarted := pipeline.GenericPipeline(ctx, orders, o.Start)
	startedToProcessed := fan.Fan(
		ctx,
		fanLimit,
		o.Process,
		createdToStarted,
	)
	processedToComplete := pipeline.GenericPipeline(ctx, startedToProcessed, o.Complete)

	return processedToComplete
}
