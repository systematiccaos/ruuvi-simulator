package model

type Tag struct {
	Sensors []Sensor `json:"sensors"`
	Address string   `json:"address"`
	Name    string   `json:"name"`
}
