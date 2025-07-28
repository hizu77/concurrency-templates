package order

import (
	"log/slog"

	"github.com/hizu77/concurrency-templates/internal/service"
)

var _ service.OrderService = (*serviceImpl)(nil)

type serviceImpl struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *serviceImpl {
	return &serviceImpl{
		logger: logger,
	}
}
