package model

type Book struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Author   string  `json:"author"`
	Price    float32 `json:"price"`
	Quantity int     `json:"available"`
}
