package storage

import (
	"context"
	"onelab/config"
	"onelab/internal/model"
	"onelab/internal/storage/postgre"
)

type IUserRepository interface {
	SignUp(ctx context.Context, user model.User) (string, error)
	GetOrders(ctx context.Context, userID string) ([]model.Order, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetByUsername(ctx context.Context, userName string) (model.User, error)
	GetSpentMoney(ctx context.Context) ([]model.SpentMoney, error)
	RentBook(ctx context.Context, username string, book model.Book) (model.Transaction, error)
	ReplenishBalance(ctx context.Context, username string, amount float32) error
}
type IBookRepository interface {
	Create(ctx context.Context, book model.Book) (string, error)
	GetAvailable(ctx context.Context) ([]model.Book, error)
	GetAllBooks(ctx context.Context) ([]model.Book, error)
	GetByID(ctx context.Context, id string) (model.Book, error)
	Update(ctx context.Context, book model.Book) error
	ReturnBook(ctx context.Context, bookId string) error
}

type IOrderRepository interface {
	Create(ctx context.Context, order model.Order) (string, error)
	GetAllOrders(ctx context.Context) ([]model.Order, error)
	GetNotReturned(ctx context.Context) ([]model.Order, error)
	GetLastMonthOrders(ctx context.Context) ([]model.Order, error)
	Update(ctx context.Context, order model.Order) error
	ReturnBook(ctx context.Context, orderId string) (model.Order, error)
}

type Storage struct {
	User  IUserRepository
	Book  IBookRepository
	Order IOrderRepository
}

func NewStorage(ctx context.Context, cfg *config.Config) (*Storage, error) {
	DB, err := postgre.Dial(ctx, cfg.Database.PgUrl)
	if err != nil {
		return nil, err
	}
	userRepo := postgre.NewUserRepository(DB)
	bookRepo := postgre.NewBookRepository(DB)
	orderRepo := postgre.NewOrderRepository(DB)

	storage := Storage{
		User:  userRepo,
		Book:  bookRepo,
		Order: orderRepo,
	}
	return &storage, nil

}
