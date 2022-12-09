package model

type StayerGetResponse struct {
	ID     int64            `json:"id"`
	Name   string           `json:"name"`
	Room   string           `json:"room"`
	RoomID int              `json:"roomID"`
	Tags   []TagGetResponse `json:"tags"`
}

type TagGetResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type LogGetResponse struct {
	ID      int64  `json:"id"`
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
	Room    string `json:"room"`
	Name    string `json:"name"`
}

type UserInformationGetResponse struct {
	ID   int64            `json:"id"`
	Name string           `json:"name"`
	Tags []TagGetResponse `json:"tags"`
}

type UserDetailInformationGetResponse struct {
	Email string `json:"email"`
	Role  int64  `json:"role"`
	UserInformationGetResponse
}

type UserRoleGetResponse struct {
	ID   int64 `json:"id"`
	Role int64 `json:"role"`
}

type RoomStayTime struct {
	Date      string     `json:"date"`
	TimeRooms []TimeRoom `json:"timeRooms"`
}

type TimeRoom struct {
	Times []int  `json:"times"`
	Name  string `json:"name"`
	ID    int    `json:"id"`
}

type SimulataneousStayLogGetResponse struct {
	ID    int64             `json:"id"`
	Date  string            `json:"date"`
	Rooms []RoomGetResponse `json:"rooms"`
}

type RoomGetResponse struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	StayTimes []StayTime `json:"stayTimes"`
}

type StayTime struct {
	ID       int64  `json:"id"`
	UserName string `json:"userName"`
	StartAt  int64  `json:"startAt"`
	EndAt    int64  `json:"endAt"`
	Color    string `json:"color"`
}

type SimultaneousStayUserGetResponse struct {
	Date  string `json:"date"`
	Names []Name `json:"names"`
}

type Name struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type RequestBody struct {
	Text string `json:"text"`
}

type RegistrationUserForm struct {
	TargetID    int64  `json:"targetID"`
	TargetEmail string `json:"targetEmail"`
	TargetName  string `json:"targetName"`
	TargetRole  int64  `json:"targetRole"`
}

type BeaconRoom struct {
	Beacons []*Beacon `json:"beacons"`
	RoomID  int64     `json:"roomID"`
}

type Beacon struct {
	Uuid string `json:"uuid" form:"uuid"`
	Rssi int64  `json:"rssi" form:"rssi"`
}
