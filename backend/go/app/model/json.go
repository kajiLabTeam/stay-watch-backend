package model

type StayerGetResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Team   string `json:"team"`
	Room   string `json:"room"`
	RoomID int    `json:"roomID"`
	Tags   []Tag  `json:"tags"`
}

type LogGetResponse struct {
	ID      string `json:"id"`
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
	Room    string `json:"room"`
	Name    string `json:"name"`
	Team    string `json:"team"`
}

type UserInformationGetResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Team string `json:"team"`
	Tags []Tag  `json:"tags"`
}

type UserRoomTimeLogGetResponse struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	RoomStayTimes []RoomStayTime `json:"roomStayTimes"`
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
