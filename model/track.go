package model

import "time"

type Track struct {
	ID     uint      `json:"id" example:"1"`
	UserID uint      `json:"user_id" example:"2"`
	TaskID uint      `json:"task_id" example:"3"`
	Start  time.Time `json:"start" example:"2006-01-02 15:04:05"`
	Stop   time.Time `json:"stop" example:"2010-01-02 15:04:05"`
}

type TrackCreate struct {
	UserID uint   `json::"user_id" example:"2"`
	TaskID uint   `json::"task_id" example:"3"`
	Start  string `json:"start" example:"2006-01-02 15:04:05"`
	Stop   string `json:"stop" example:"2010-01-02 15:04:05"`
}

type TrackUpdate struct {
	ID     uint    `json:"id" example:"1"`
	UserID *uint   `json::"user_id" example:"2"`
	TaskID *uint   `json::"task_id" example:"3"`
	Start  *string `json:"start" example:"2006-01-02 15:04:05"`
	Stop   *string `json:"stop" example:"2010-01-02 15:04:05"`
}

type TrackDelete struct {
	ID uint `json:"id" example:"1"`
}
