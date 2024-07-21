package model

type User struct {
	ID         uint
	Name       string
	Surname    string
	Patronymic *string
	Address    string
	Passport   int
}

type UserCreate struct {
	Name       string
	Surname    string
	Patronymic *string
	Address    string
	Passport   int
}

type UserUpdate struct {
	ID         uint
	Name       *string
	Surname    *string
	Patronymic *string
	Address    *string
	Passport   *int
}

type UserDelete struct {
	ID uint
}
