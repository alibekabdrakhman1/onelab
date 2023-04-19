package service

import "onelab/internal/storage"

type Service struct {
	Book  IBookService
	User  IUserService
	Order IOrderService
}

func NewManager(storage *storage.Storage) (*Service, error) {
	bookService := NewBookService(storage)
	userService := NewUserService(storage)
	orderService := NewOrderService(storage)

	return &Service{
		Book:  bookService,
		User:  userService,
		Order: orderService,
	}, nil
}
