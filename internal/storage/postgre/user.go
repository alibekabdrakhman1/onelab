package postgre

import (
	"context"
	"gorm.io/gorm"
	"onelab/internal/model"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) SignUp(ctx context.Context, user model.User) (string, error) {
	id := user.ID
	if err := r.DB.Create(&user).Error; err != nil {
		return "", err
	}
	return id, nil
}

//func (r *UserRepository) GetBooks(ctx context.Context, userID string) ([]model.Order, error) {
//
//}

// ya znauy chto tak delat nelzya
func (r *UserRepository) GetSpentMoney(ctx context.Context) ([]model.SpentMoney, error) {
	users, _ := r.GetAllUsers(ctx)
	ans := make([]model.SpentMoney, len(users))

	for i := 0; i < len(users); i++ {
		ans[i].Login = users[i].Login
		ans[i].Money = users[i].SpentMoney
	}
	return ans, nil
}

func (r *UserRepository) RentBook(ctx context.Context, username string, book model.Book) (model.Transaction, error) {
	var res model.User
	err := r.DB.WithContext(ctx).Where("login = ?", username).Find(&res).Error
	if err != nil {
		return model.Transaction{}, err
	}
	res.Balance -= book.Price
	err = r.DB.WithContext(ctx).Where("login = ?", username).Updates(&res).Error
	return model.Transaction{Username: username, TypeOfTransaction: "-", Amount: book.Price, Description: "Rent Book"}, err
}

func (r *UserRepository) ReplenishBalance(ctx context.Context, username string, amount float32) error {
	var res model.User
	err := r.DB.WithContext(ctx).Where("login = ?", username).Find(&res).Error
	if err != nil {
		return err
	}
	res.Balance += amount
	err = r.DB.WithContext(ctx).Where("login = ?", username).Updates(&res).Error
	return err
}

func (r *UserRepository) GetOrders(ctx context.Context, userID string) ([]model.Order, error) {
	var orders []model.Order
	if err := r.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
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
