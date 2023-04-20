package service

import (
	"context"
	"github.com/google/uuid"
	"onelab/internal/model"
	"onelab/internal/storage"
)

type IBookService interface {
	Create(ctx context.Context, book model.Book) (string, error)
	GetAvailable(ctx context.Context) ([]model.Book, error)
	GetAllBooks(ctx context.Context) ([]model.Book, error)
	GetByID(ctx context.Context, id string) (model.Book, error)
	Update(ctx context.Context, book model.Book) error
}

type BookService struct {
	repository *storage.Storage
}

func (s *BookService) Update(ctx context.Context, book model.Book) error {
	return s.repository.Book.Update(ctx, book)
}

func NewBookService(r *storage.Storage) *BookService {
	return &BookService{
		repository: r,
	}
}
func (s *BookService) Create(ctx context.Context, book model.Book) (string, error) {
	book.ID = uuid.NewString()
	return s.repository.Book.Create(ctx, book)
}

func (s *BookService) GetAvailable(ctx context.Context) ([]model.Book, error) {
	return s.repository.Book.GetAvailable(ctx)
}

func (s *BookService) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	return s.repository.Book.GetAllBooks(ctx)
}

func (s *BookService) GetByID(ctx context.Context, id string) (model.Book, error) {
	return s.repository.Book.GetByID(ctx, id)
}
