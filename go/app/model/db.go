package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UUID  string
	Name  string
	Email string
	Role  int64
}

type Log struct {
	gorm.Model
	RoomID  int64
	StartAt time.Time
	EndAt   time.Time
	UserID  int64
	Rssi    int64
}

type Room struct {
	gorm.Model
	Name string
}

type Stayer struct {
	gorm.Model
	UserID int64
	RoomID int64
	Rssi   int64
}

type Tag struct {
	gorm.Model
	Name string
}

type TagMap struct {
	gorm.Model
	UserID int64
	TagID  int64
}

type Attendance struct {
	gorm.Model
	UserID int64
	Date   string
	Flag   bool
}

type AttendanceTmp struct {
	UserID int64 `json:"userID" xorm:"user_id"`
	Flag   int64 `json:"exit" xorm:"flag"`
}
