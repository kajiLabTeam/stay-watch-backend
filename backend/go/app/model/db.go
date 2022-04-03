package model

type User struct {
	ID   string `json:"ID" xorm:"id"`
	Name string `json:"name" xorm:"name"`
	Team string `json:"team" xorm:"team"`
}

type Log struct {
	RoomID  int64  `json:"roomID" xorm:"room_id"`
	StartAt string `json:"startAt" xorm:"start_at"`
	EndAt   string `json:"endAt" xorm:"end_at"`
	UserID  string `json:"userID" xorm:"user_id"`
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
	UserID string `json:"userID" xorm:"user_id"`
	RoomID int64  `json:"roomID" xorm:"room_id"`
	Rssi   int64  `json:"rssi" xorm:"rssi"`
}
