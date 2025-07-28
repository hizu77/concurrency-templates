package order

import (
	"context"
	"log/slog"
	"time"

	model "github.com/hizu77/concurrency-templates/internal/model/order"
)

func (o *serviceImpl) Start(ctx context.Context, order model.Order) model.Order {
	time.Sleep(time.Second)

	if err := ctx.Err(); err != nil {
		return model.Order{}
	}

	order.State = model.StateStarted

	o.logger.Info("order started", slog.Any("order_id", order.ID))

	return order
}
