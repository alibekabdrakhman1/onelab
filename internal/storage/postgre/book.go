package postgre

import (
	"context"
	"gorm.io/gorm"
	"onelab/internal/model"
)

type BookRepository struct {
	DB *gorm.DB
}

func (r *BookRepository) Create(ctx context.Context, user model.Book) (string, error) {
	id := user.ID
	if err := r.DB.Create(&user).Error; err != nil {
		return "", err
	}
	return id, nil
}
func (r *BookRepository) GetAvailable(ctx context.Context) ([]model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *BookRepository) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *BookRepository) GetByAuthor(ctx context.Context, author string) (model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		DB: db,
	}
}
