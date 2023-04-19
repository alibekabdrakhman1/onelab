package model

type Book struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	Available bool   `json:"available"`
}
