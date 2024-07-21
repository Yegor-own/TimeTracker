package model

import "time"

type Track struct {
	ID     uint
	UserID uint
	TaskID uint
	Start  time.Time
	Stop   *time.Time
}

type TrackCreate struct {
	UserID uint
	TaskID uint
	Start  time.Time
}

type TrackUpdate struct {
	ID     uint
	UserID *uint
	TaskID *uint
	Start  *time.Time
	Stop   *time.Time
}

type TrackDelete struct {
	ID uint
}
