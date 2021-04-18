package domain

type Game struct {
	Id         int        `json:"game_index"`
	Generation Generation `json:"generation"`
}
