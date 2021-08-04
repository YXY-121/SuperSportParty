package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type A struct {
	Type string
	Data string
}
type B struct {
	SenderId string
	AccepterId string
}
type ServerController struct {
	beego.Controller
}

func (i* ServerController)Post()   {


	a:=A{
		Type: "A",
		Data: "HAHA",
	}

	fmt.Println(json.Marshal(a))
	me,_:=json.Marshal(a)
	b:=A{}
	json.Unmarshal(me, &b)
	fmt.Println("type",b.Type)
	fmt.Println("data",b.Data)


}
