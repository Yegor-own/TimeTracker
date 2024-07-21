package model

type Task struct {
	ID          uint
	Name        string
	Description string
}

type TaskCreate struct {
	Name        string
	Description string
}

type TaskUpdate struct {
	ID          uint
	Name        *string
	Description *string
}

type TaskDelete struct {
	ID uint
}
