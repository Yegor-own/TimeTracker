package model

type User struct {
	ID         uint
	Name       string
	Surname    string
	Patronymic *string
	Address    string
}

type UserCreate struct {
	Name       string
	Surname    string
	Patronymic *string
	Address    string
}

type UserUpdate struct {
	UserID     uint
	Name       *string
	Surname    *string
	Patronymic *string
	Address    *string
}

type UserDelete struct {
	UserID uint
}
