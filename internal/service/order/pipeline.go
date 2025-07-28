package order

import (
	"context"

	"github.com/hizu77/concurrency-templates/internal/infra/pipeline"
	model "github.com/hizu77/concurrency-templates/internal/model/order"
)

func (o *serviceImpl) Pipeline(ctx context.Context, orders <-chan model.Order) <-chan model.Order {
	createdToStarted := pipeline.GenericPipeline(
		ctx,
		orders,
		o.Start,
	)

	startedToProcessed := pipeline.GenericPipeline(
		ctx,
		createdToStarted,
		o.Process,
	)

	processedToCompleted := pipeline.GenericPipeline(
		ctx,
		startedToProcessed,
		o.Complete,
	)

	return processedToCompleted
}
