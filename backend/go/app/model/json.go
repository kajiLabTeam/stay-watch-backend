package model

type StayerGetResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Team string `json:"team"`
	Room string `json:"room"`
}

type LogGetResponse struct {
	ID      string `json:"id"`
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
	Room    string `json:"room"`
	Name    string `json:"name"`
	Team    string `json:"team"`
}
