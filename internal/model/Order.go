package model

import "time"

type Order struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	BookID      string    `json:"book_id" gorm:"references:ID"`
	UserID      string    `json:"user_id" gorm:"references:ID"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserBooks struct {
	User  User
	Books []Book
}

type Transaction struct {
	Username          string  `json:"username"`
	TypeOfTransaction string  `json:"type_of_transaction"`
	Amount            float32 `json:"amount"`
	Description       string  `json:"description"`
}
