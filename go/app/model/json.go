package model

type StayerGetResponse struct {
	ID     int64            `json:"id"`
	Name   string           `json:"name"`
	Room   string           `json:"room"`
	RoomID int              `json:"roomId"`
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

type LogWithCount struct {
	Logs  []LogGetResponse `json:"logs"`
	Count int              `json:"count"`
}

type UserInformationGetResponse struct {
	ID   int64            `json:"id"`
	Name string           `json:"name"`
	Tags []TagGetResponse `json:"tags"`
}

type ExtendedUserInformationGetResponse struct {
	ID    int64            `json:"id"`
	Name  string           `json:"name"`
	Tags  []TagGetResponse `json:"tags"`
	Uuid  string           `json:"uuid"`
	Email string           `json:"email"`
	Role  int64            `json:"role"`
}

type UserRoleCommunityGetResponse struct {
	ID            int64  `json:"id"`
	UUID          string `json:"uuid"`
	Name          string `json:"name"`
	Role          int64  `json:"role"`
	CommunityId   int64  `json:"communityId"`
	CommunityName string `json:"communityName"`
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
	ID    int64  `form:"id"`
	Email string `form:"email"`
	Name  string `form:"name"`
	Role  int64  `form:"role"`
}

// バックエンドからフロントへ返すユーザ情報
type UserEditorResponse struct {
	ID                 int64            `json:"id"`
	Name               string           `json:"name"`
	Uuid               string           `json:"uuid"`
	Email              string           `json:"email"`
	Role               int64            `json:"role"`
	BeaconUuidEditable bool             `json:"beaconUuidEditable"`
	BeaconName         string           `json:"beaconName"`
	Tags               []TagGetResponse `json:"tags"`
}

// フロントからバックエンドへ送られてきた新規作成ユーザ情報
type UserCreateRequest struct {
	Name        string  `json:"name"`
	Uuid        string  `json:"uuid"`
	Email       string  `json:"email"`
	Role        int64   `json:"role"`
	CommunityId int64   `json:"communityId"`
	BeaconName  string  `json:"beaconName"`
	TagIds      []int64 `json:"tagIds"`
	PrivateKey  string  `json:"privateKey"`
}

// フロントからバックエンドへ送られてきた更新するユーザ情報
type UserUpdateRequest struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Uuid        *string `json:"uuid"`
	Email       *string `json:"email"`
	Role        *int64  `json:"role"`
	CommunityId *int64  `json:"communityId"`
	BeaconName  string  `json:"beaconName"`
	TagIds      []int64 `json:"tagIds"`
	PrivateKey  *string `json:"privateKey"`
}

type BeaconRoom struct {
	Beacons []*BeaconSignal `json:"beacons"`
	RoomID  int64           `json:"roomId"`
}

type BeaconSignal struct {
	Uuid string `json:"uuid" form:"uuid"`
	Rssi int64  `json:"rssi" form:"rssi"`
}

type BeaconGetResponse struct {
	BeaconId     int64  `json:"beaconId"`
	BeaconName   string `json:"beaconName"`
	UuidEditable bool   `json:"uuidEditable"`
}

type RoomEditorForm struct {
	RoomID     int64     `json:"roomId"`
	RoomName   string    `json:"roomName"`
	Polygon    [][]int64 `json:"polygon"`
	BuildingID int64     `json:"buildingId"`
}

type RoomsGetResponse struct {
	RoomID        int64     `json:"roomId"`
	Name          string    `json:"roomName"`
	CommunityName string    `json:"communityName"`
	BuildingName  string    `json:"buildingName"`
	Polygon       [][]int64 `json:"polygon"`
	BuildingId    int64     `json:"buildingId"`
}

type BuildingsEditorGetResponse struct {
	BuildingID   int64  `json:"buildingId"`
	Name         string `json:"buildingName"`
	MapImagePath string `json:"buildingImagePath"`
}

type CommunityGetResponse struct {
	CommunityId int64  `json:"id"`
	Name        string `json:"name"`
}

type TagsNamesGetResponse struct {
	Name string `json:"tagName"`
}

type TagsGetResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
