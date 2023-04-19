package handler

import (
	"onelab/internal/service"
	"onelab/internal/transport/http/middleware"
)

type Manager struct {
	Book  IBookHandler
	Order IOrderHandler
	User  IUserHandler
}

func NewManager(srv *service.Service, jwt *middleware.JWTAuth) *Manager {
	return &Manager{
		Book:  NewBookHandler(srv),
		Order: NewOrderHandler(srv),
		User:  NewUserHandler(srv, jwt),
	}
}
