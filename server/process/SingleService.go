package process

import (
	"apiproject/common"
	"encoding/json"
	"fmt"
)

type SingleService struct {

}
func NewSingleService() *SingleService{
	return &SingleService{
	}
}

func (s *SingleService)SendSingleMessage(data string)  {
	singleMessage:=common.SingleMessage{}
	fmt.Println(json.Unmarshal([]byte(data),&singleMessage))

}