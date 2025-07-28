package service

import (
	"context"

	model "github.com/hizu77/concurrency-templates/internal/model/order"
)

type (
	OrderService interface {
		Start(ctx context.Context, order model.Order) (model.Order, error)
		Process(ctx context.Context, order model.Order) (model.Order, error)
		Complete(ctx context.Context, order model.Order) (model.Order, error)
	}
)
