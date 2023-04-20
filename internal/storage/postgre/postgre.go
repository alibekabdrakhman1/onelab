package postgre

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"onelab/internal/model"
)

func Dial(ctx context.Context, url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	if db != nil {
		err := db.WithContext(ctx).AutoMigrate(&model.User{}, &model.Book{}, &model.Order{})
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
