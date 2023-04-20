package postgre

import (
	"context"
	"gorm.io/gorm"
	"onelab/internal/model"
)

type BookRepository struct {
	DB *gorm.DB
}

func (r *BookRepository) ReturnBook(ctx context.Context, bookId string) error {
	var res model.Book
	err := r.DB.WithContext(ctx).Where("id = ?", bookId).Find(&res).Error
	if err != nil {
		return err
	}
	res.Quantity++
	err = r.DB.WithContext(ctx).Where("id = ?", bookId).Updates(&res).Error
	return err
}

func (r *BookRepository) Update(ctx context.Context, book model.Book) error {
	//TODO implement me
	panic("implement me")
}

func (r *BookRepository) Create(ctx context.Context, user model.Book) (string, error) {
	id := user.ID
	if err := r.DB.Create(&user).Error; err != nil {
		return "", err
	}
	return id, nil
}
func (r *BookRepository) GetAvailable(ctx context.Context) ([]model.Book, error) {
	var resp []model.Book
	err := r.DB.WithContext(ctx).Where("quantity > 0").Find(&resp).Error
	return resp, err
}

func (r *BookRepository) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	var res []model.Book
	err := r.DB.WithContext(ctx).Find(&res).Error
	return res, err
}

func (r *BookRepository) GetByID(ctx context.Context, id string) (model.Book, error) {
	var res model.Book
	err := r.DB.WithContext(ctx).Where("id = ?", id).Find(&res).Error
	return res, err
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		DB: db,
	}
}
