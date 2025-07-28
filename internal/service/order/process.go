package order

import (
	"context"
	"log/slog"
	"time"

	model "github.com/hizu77/concurrency-templates/internal/model/order"
)

func (o *serviceImpl) Process(ctx context.Context, order model.Order) (model.Order, error) {
	time.Sleep(time.Second * 3)

	if err := ctx.Err(); err != nil {
		return model.Order{}, nil
	}

	order.State = model.StateProcessed

	o.logger.Info("order processed", slog.Any("order_id", order.ID))

	return order, nil
}
