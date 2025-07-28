package service

import (
	"context"

	model "github.com/hizu77/concurrency-templates/internal/model/order"
)

type (
	OrderService interface {
		Start(ctx context.Context, order model.Order) model.Order
		Process(ctx context.Context, order model.Order) model.Order
		Complete(ctx context.Context, order model.Order) model.Order
		Pipeline(ctx context.Context, orders <-chan model.Order) <-chan model.Order
		FanPipeline(ctx context.Context, orders <-chan model.Order, fanLimit int) <-chan model.Order
	}
)
