package postgre

import (
	"context"
	"gorm.io/gorm"
	"onelab/internal/model"
	"time"
)

type OrderRepository struct {
	DB *gorm.DB
}

func (r *OrderRepository) Create(ctx context.Context, user model.Order) (string, error) {
	id := user.ID
	if err := r.DB.Create(&user).Error; err != nil {
		return "", err
	}
	return id, nil
}

func (r *OrderRepository) GetAllOrders(ctx context.Context) ([]model.Order, error) {
	var resp []model.Order
	err := r.DB.WithContext(ctx).Find(&resp)
	return resp, err.Error
}

func (r *OrderRepository) GetNotReturned(ctx context.Context) ([]model.Order, error) {
	var resp []model.Order
	err := r.DB.WithContext(ctx).Where("returned = ", false).Find(&resp).Error
	return resp, err
}

func (r *OrderRepository) GetLastMonthOrders(ctx context.Context) ([]model.Order, error) {
	var resp []model.Order
	err := r.DB.WithContext(ctx).Where("ordered_date >= NOW() - INTERVAL '1 MONTH'").Find(&resp).Error
	return resp, err
}
func (r *OrderRepository) ReturnBook(ctx context.Context, orderId string) error {
	var record model.Order
	if err := r.DB.WithContext(ctx).Model(&model.Order{}).Where("id = ?", orderId).First(&record).Error; err != nil {
		return err
	}
	record.Returned = true
	record.ReturnedDate = time.Now()
	if err := r.DB.WithContext(ctx).Model(&record).Updates(&record).Error; err != nil {
		return err
	}
	return nil
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}
