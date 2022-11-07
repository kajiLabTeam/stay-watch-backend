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
	ID   int64  `json:"id" xorm:"id"`
	Name string `json:"name" xorm:"name"`
}

type TagMap struct {
	ID     int64  `json:"ID" xorm:"id"`
	UserID int64  `json:"userID" xorm:"user_id"`
	TagID  string `json:"tagID" xorm:"tag_id"`
}

type Attendance struct {
	ID     int64  `json:"ID" xorm:"id"`
	UserID int64  `json:"userID" xorm:"user_id"`
	Date   string `json:"date" xorm:"date"`
	Flag   bool   `json:"exit" xorm:"flag"`
}

type AttendanceTmp struct {
	UserID int64 `json:"userID" xorm:"user_id"`
	Flag   int64 `json:"exit" xorm:"flag"`
}
