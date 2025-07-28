package order

import (
	"context"
	"log/slog"

	model "github.com/hizu77/concurrency-templates/internal/model/order"
)

func (o *serviceImpl) Complete(ctx context.Context, order model.Order) model.Order {
	if err := ctx.Err(); err != nil {
		return model.Order{}
	}

	order.State = model.StateCompleted

	o.logger.Info("order completed", slog.Any("order_id", order.ID))

	return order
}
