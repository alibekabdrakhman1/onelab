package model

import "time"

type Order struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	BookID       string    `json:"book_id" gorm:"references:ID"`
	Book         Book      `gorm:"ForeignKey:BookID"`
	UserID       string    `json:"user_id" gorm:"references:ID"`
	User         User      `gorm:"ForeignKey:UserID"`
	Returned     bool      `json:"returned"`
	OrderedDate  time.Time `json:"ordered_date"`
	ReturnedDate time.Time `json:"returned_date"`
}
