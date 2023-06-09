package service

import (
	"context"
	"github.com/google/uuid"
	"onelab/internal/model"
	"onelab/internal/storage"
	"time"
)

type IOrderService interface {
	Create(ctx context.Context, order model.Order) (string, error)
	GetAllOrders(ctx context.Context) ([]model.Order, error)
	GetNotReturned(ctx context.Context) ([]model.Order, error)
	GetLastMonthOrders(ctx context.Context) ([]model.Order, error)
	Update(ctx context.Context, order model.Order) error
}
type OrderService struct {
	repository *storage.Storage
}

func (s *OrderService) Update(ctx context.Context, order model.Order) error {
	return s.repository.Order.Update(ctx, order)
}

func NewOrderService(r *storage.Storage) *OrderService {
	return &OrderService{
		repository: r,
	}
}
func (s *OrderService) Create(ctx context.Context, order model.Order) (string, error) {
	order.ID = uuid.NewString()
	order.UpdatedAt = time.Now()
	return s.repository.Order.Create(ctx, order)
}

func (s *OrderService) GetAllOrders(ctx context.Context) ([]model.Order, error) {
	return s.repository.Order.GetAllOrders(ctx)
}

func (s *OrderService) GetNotReturned(ctx context.Context) ([]model.Order, error) {
	return s.repository.Order.GetNotReturned(ctx)
}

func (s *OrderService) GetLastMonthOrders(ctx context.Context) ([]model.Order, error) {
	return s.repository.Order.GetLastMonthOrders(ctx)
}
