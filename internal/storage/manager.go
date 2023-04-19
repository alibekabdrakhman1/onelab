package storage

import (
	"context"
	"onelab/config"
	"onelab/internal/model"
	"onelab/internal/storage/postgre"
)

type IUserRepository interface {
	GetOrders(ctx context.Context, userID string) ([]model.Order, error)
	SignUp(ctx context.Context, user model.User) (string, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetByUsername(ctx context.Context, userName string) (model.User, error)
}
type IBookRepository interface {
	Create(ctx context.Context, book model.Book) (string, error)
	GetAvailable(ctx context.Context) ([]model.Book, error)
	GetAllBooks(ctx context.Context) ([]model.Book, error)
	GetByAuthor(ctx context.Context, author string) (model.Book, error)
}

type IOrderRepository interface {
	Create(ctx context.Context, order model.Order) (string, error)
	ReturnBook(ctx context.Context, orderId string) error
	GetAllOrders(ctx context.Context) ([]model.Order, error)
	GetNotReturned(ctx context.Context) ([]model.Order, error)
	GetLastMonthOrders(ctx context.Context) ([]model.Order, error)
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
