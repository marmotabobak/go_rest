package models

type Item struct {
	Data ItemData `json:"data"`
}

type ItemData struct {
	Value string `json:"value"`
}