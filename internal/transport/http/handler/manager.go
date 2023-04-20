package handler

import (
	"onelab/internal/service"
	"onelab/internal/transport/http/middleware"
	transactions "onelab/proto"
)

type Manager struct {
	Book        IBookHandler
	Order       IOrderHandler
	User        IUserHandler
	Transaction ITransactionHandler
}

func NewManager(srv *service.Service, jwt *middleware.JWTAuth, grpc transactions.TransactionServiceClient) *Manager {
	return &Manager{
		Book:        NewBookHandler(srv),
		Order:       NewOrderHandler(srv),
		User:        NewUserHandler(srv, jwt),
		Transaction: NewTransactionHandler(srv, grpc),
	}
}
