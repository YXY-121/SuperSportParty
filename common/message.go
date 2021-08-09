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
type CreateGroupMessage struct {
	GroupName string `json:"group_name"`
	GroupId string 	`json:"group_id"`
	UserIds []string `json:"user_ids"`
	CreaterId string `json:"creater_id"`
}

type RedPaper struct {
	paperId string
	SenderId string
	Money string
	Type string
}
const GroupMessageType = "group"
const SingleMessageType = "single"
const CreateGroupType ="createGroup"
