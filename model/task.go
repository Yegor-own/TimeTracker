package model

type Task struct {
	ID          uint   `json:"id" example:"1"`
	Name        string `json:"name" example:"code review"`
	Description string `json:"description" example:"code review by lead"`
}

type TaskCreate struct {
	Name        string `json:"name" example:"code review"`
	Description string `json:"description" example:"code review by lead"`
}

type TaskUpdate struct {
	ID          uint    `json:"id" example:"1"`
	Name        *string `json:"name" example:"code review"`
	Description *string `json:"description" example:"code review by lead"`
}

type TaskDelete struct {
	ID uint `json:"id" example:"1"`
}
