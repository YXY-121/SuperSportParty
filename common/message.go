package common

type GroupMessage struct {
	SenderId string `json:"sender_id"`
	Data string` json:"data"`
	Type string `json:"type"`
	GroupId string `json:"group_id"`
}
type SingleMessage struct {
	SenderId string `json:"sender_id"`
	Data string `json:"data"`
	Type string `json:"type"`
	AccepterId string `json:"accepter_id"`
}

type Message struct {
	Type string `json:"type"` //single或者是group，作为
	Data string `json:"data"`
}