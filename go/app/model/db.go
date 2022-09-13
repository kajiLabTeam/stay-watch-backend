package model

type User struct {
	ID    int64  `json:"ID" xorm:"id"`
	UUID  string `json:"uuid" xorm:"uid"`
	Name  string `json:"name" xorm:"name"`
	Email string `json:"email" xorm:"email"`
	Role  int64  `json:"role" xorm:"role"`
}

type Log struct {
	ID      int64  `json:"ID" xorm:"id"`
	RoomID  int64  `json:"roomID" xorm:"room_id"`
	StartAt string `json:"startAt" xorm:"start_at"`
	EndAt   string `json:"endAt" xorm:"end_at"`
	UserID  int64  `json:"userID" xorm:"user_id"`
	Rssi    int64  `json:"rssi" xorm:"rssi"`
}

type BeaconRoom struct {
	Beacons []*Beacon `json:"beacons"`
	RoomID  int64     `json:"roomID"`
}

type Beacon struct {
	Uuid string `json:"uuid" form:"uuid"`
	Rssi int64  `json:"rssi" form:"rssi"`
}

type Room struct {
	ID   int64  `json:"ID" xorm:"id"`
	Name string `json:"name" xorm:"name"`
}

type Stayer struct {
	UserID int64 `json:"userID" xorm:"user_id"`
	RoomID int64 `json:"roomID" xorm:"room_id"`
	Rssi   int64 `json:"rssi" xorm:"rssi"`
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
