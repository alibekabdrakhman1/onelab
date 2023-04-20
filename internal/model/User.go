package model

type User struct {
	ID         string  `json:"id" db:"id"`
	Name       string  `json:"name" validate:"required"`
	Surname    string  `json:"surname" validate:"required"`
	Login      string  `json:"login" validate:"required"`
	Password   string  `json:"password" validate:"required"`
	Balance    float32 `json:"balance"`
	SpentMoney float32 `json:"spent_money"`
}

type LogIn struct {
	Login    string `validate:"required"`
	Password string `validate:"required"`
}

type SpentMoney struct {
	Login string  `json:"login"`
	Money float32 `json:"money"`
}
