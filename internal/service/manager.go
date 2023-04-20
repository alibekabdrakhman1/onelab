package service

import (
	"onelab/internal/storage"
	transactions "onelab/proto"
)

type Service struct {
	Book  IBookService
	User  IUserService
	Order IOrderService
}

func NewManager(storage *storage.Storage, grpc transactions.TransactionServiceClient) (*Service, error) {
	bookService := NewBookService(storage)
	userService := NewUserService(storage, grpc)
	orderService := NewOrderService(storage)

	return &Service{
		Book:  bookService,
		User:  userService,
		Order: orderService,
	}, nil
}
