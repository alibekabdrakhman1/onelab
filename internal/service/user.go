package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"onelab/internal/model"
	"onelab/internal/storage"
	transactions "onelab/proto"
)

type IUserService interface {
	Login(ctx context.Context, user model.LogIn) (*model.ContextData, error)
	SignUp(ctx context.Context, user model.User) (string, error)
	GetOrders(ctx context.Context, userID string) ([]model.Order, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetByUsername(ctx context.Context, userName string) (model.User, error)
	GetBooks(ctx context.Context, userID string) ([]model.Order, error)
	GetSpentMoney(ctx context.Context) ([]model.SpentMoney, error)
	RentBook(ctx context.Context, username string, bookId string) (model.Transaction, error)
	ReturnBook(ctx context.Context, orderId string) error
	ReplenishBalance(ctx context.Context, username string, amount float32) (string, error)
}

type UserService struct {
	repository *storage.Storage
	grpc       transactions.TransactionServiceClient
}

func (s *UserService) GetBooks(ctx context.Context, userID string) ([]model.Order, error) {
	//return s.repository.User.GetBooks(ctx, userID)
	panic("zhok")
}

func (s *UserService) GetSpentMoney(ctx context.Context) ([]model.SpentMoney, error) {
	return s.repository.User.GetSpentMoney(ctx)
}

func (s *UserService) RentBook(ctx context.Context, username string, bookId string) (model.Transaction, error) {
	book, err := s.repository.Book.GetByID(ctx, bookId)
	if err != nil {
		return model.Transaction{}, err
	}
	user, err := s.repository.User.GetByUsername(ctx, username)
	if err != nil {
		return model.Transaction{}, err
	}
	transaction := &transactions.CreateTransRequest{Transaction: &transactions.Transaction{
		Username:    username,
		Type:        "-",
		Amount:      int32(book.Price),
		Description: "Rent Book",
	}}
	// tak tozhe nelzya
	_, err = s.grpc.Create(ctx, transaction)
	if err != nil {
		return model.Transaction{}, err
	}
	_, err = s.repository.Order.Create(ctx, model.Order{BookID: bookId, UserID: user.ID, TotalAmount: float64(transaction.Transaction.Amount)})
	if err != nil {
		return model.Transaction{}, err
	}

	if err != nil {
		return model.Transaction{}, err
	}
	return model.Transaction{Username: transaction.Transaction.Username, TypeOfTransaction: transaction.Transaction.Type, Amount: float32(transaction.Transaction.Amount), Description: transaction.Transaction.Description}, nil
}

func (s *UserService) ReturnBook(ctx context.Context, orderId string) error {
	order, err := s.repository.Order.ReturnBook(ctx, orderId)
	if err != nil {
		return err
	}
	err = s.repository.Book.ReturnBook(ctx, order.BookID)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) ReplenishBalance(ctx context.Context, username string, amount float32) (string, error) {
	t := &transactions.CreateTransRequest{Transaction: &transactions.Transaction{
		Username:    username,
		Type:        "+",
		Amount:      int32(amount),
		Description: "Popolnenie",
	}}
	err := s.repository.User.ReplenishBalance(ctx, username, amount)
	if err != nil {
		return "", err
	}
	tr, err := s.grpc.Create(ctx, t)

	return tr.Id, err
}

func NewUserService(r *storage.Storage, grpc transactions.TransactionServiceClient) *UserService {
	return &UserService{
		repository: r,
		grpc:       grpc,
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
