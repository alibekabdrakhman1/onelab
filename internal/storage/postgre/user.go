package postgre

import (
	"context"
	"gorm.io/gorm"
	"onelab/internal/model"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) GetOrders(ctx context.Context, userID string) ([]model.Order, error) {

	var orders []model.Order
	if err := r.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *UserRepository) SignUp(ctx context.Context, user model.User) (string, error) {
	id := user.ID
	if err := r.DB.Create(&user).Error; err != nil {
		return "", err
	}
	return id, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var resp []model.User
	err := r.DB.WithContext(ctx).Find(&resp)
	return resp, err.Error
}

func (r *UserRepository) GetByUsername(ctx context.Context, userName string) (model.User, error) {
	var res model.User
	err := r.DB.WithContext(ctx).Where("login = ?", userName).Find(&res).Error
	return res, err
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}
