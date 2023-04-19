package model

type User struct {
	ID       string `json:"id" db:"id"`
	Name     string `json:"name" validate:"required"`
	Surname  string `json:"surname" validate:"required"`
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LogIn struct {
	Login    string `validate:"required"`
	Password string `validate:"required"`
}

type UpdatePassword struct {
	Login       string
	OldPassword string
	NewPassword string
}
