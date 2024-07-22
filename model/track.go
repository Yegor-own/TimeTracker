package model

import "time"

type Track struct {
	ID     uint
	UserID uint
	TaskID uint
	Start  time.Time
	Stop   time.Time
}

type TrackCreate struct {
	UserID uint `json:"user_id"`
	TaskID uint `json:"task_id"`
	Start  string
	Stop   string
}

type TrackUpdate struct {
	ID     uint
	UserID *uint `json:"user_id"`
	TaskID *uint `json:"task_id"`
	Start  *string
	Stop   *string
}

type TrackDelete struct {
	ID uint
}
