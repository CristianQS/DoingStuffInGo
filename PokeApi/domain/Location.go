package domain

type Location struct {
	Id     int    `json:"ìd"`
	Name   string `json:"name"`
	Region Region `json:"region"`
	Names  []Name `json:"names"`
	Games  []Game `json:"game_indices"`
	Areas  []Area `json:"areas"`
}
