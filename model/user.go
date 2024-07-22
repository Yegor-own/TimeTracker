package model

type User struct {
	ID         uint    `json:"id" example:"1"`
	Name       string  `json:"name" example:"Jon"`
	Surname    string  `json:"surname" example:"Doe"`
	Patronymic *string `json:"patronymic" example:"White"`
	Address    string  `json:"address" example:"LA"`
	Passport   int     `json:"passport" example:"123456"`
}

type UserCreate struct {
	Name       string  `json:"name" example:"Jon"`
	Surname    string  `json:"surname" example:"Doe"`
	Patronymic *string `json:"patronymic" example:"White"`
	Address    string  `json:"address" example:"LA"`
	Passport   int     `json:"passport" example:"123456"`
}

type UserUpdate struct {
	ID         uint    `json:"id" example:"1"`
	Name       *string `json:"name" example:"Jon"`
	Surname    *string `json:"surname" example:"Doe"`
	Patronymic *string `json:"patronymic" example:"White"`
	Address    *string `json:"address" example:"LA"`
	Passport   *int    `json:"passport" example:"123456"`
}

type UserDelete struct {
	ID uint `json:"id" example:"1"`
}
