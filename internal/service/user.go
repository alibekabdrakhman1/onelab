package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"onelab/internal/model"
	"onelab/internal/storage"
)

type IUserService interface {
	GetOrders(ctx context.Context, userID string) ([]model.Order, error)
	SignUp(ctx context.Context, user model.User) (string, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetByUsername(ctx context.Context, userName string) (model.User, error)
	Login(ctx context.Context, user model.LogIn) (*model.ContextData, error)
}

type UserService struct {
	repository *storage.Storage
}

func NewUserService(r *storage.Storage) *UserService {
	return &UserService{
		repository: r,
	}
}

func (s *UserService) GetOrders(ctx context.Context, userId string) ([]model.Order, error) {
	return s.repository.User.GetOrders(ctx, userId)
}

func (s *UserService) SignUp(ctx context.Context, user model.User) (string, error) {
	user.ID = uuid.NewString()
	user.Password = generatePasswordHash(user.Password)
	return s.repository.User.SignUp(ctx, user)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.repository.User.GetAllUsers(ctx)
}

func (s *UserService) GetByUsername(ctx context.Context, userName string) (model.User, error) {
	return s.repository.User.GetByUsername(ctx, userName)
}

func (s *UserService) Login(ctx context.Context, user model.LogIn) (*model.ContextData, error) {
	user1, err := s.repository.User.GetByUsername(ctx, user.Login)
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	if err := bcrypt.CompareHashAndPassword([]byte(user1.Password), []byte(user.Password)); err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &model.ContextData{
		UserID: user1.ID,
	}, nil
}

func generatePasswordHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes)
}
