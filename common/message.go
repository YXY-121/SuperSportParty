package common

type GroupMessage struct {
	SenderId string `json:"sender_id"`
	Content string `json:"content"`
	Type string `json:"type"`
	GroupId string `json:"group_id"`
}
type SingleMessage struct {
	SenderId string `json:"sender_id"`
	Content string `json:"content"`
	Type string `json:"type"`
	AccepterId string `json:"accepter_id"`

}

type Message struct {
	GroupId string `json:"group_id"`
	UserId string	`json:"user_id"`
	Type string `json:"type"` //single或者是group，作为
	Content string `json:"content"`
}

const GroupMessageType = "group"
const SingleMessageType = "single"
