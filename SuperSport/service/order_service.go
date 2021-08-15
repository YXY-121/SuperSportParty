package service

import "apiproject/redPaper/model"

//
type OrderService struct {

}


//根据位置获取拼团活动
func GetOrderByLocation (lat1,lng1 float64)  {

}
//根据位置、运动类型获取拼团活动
func GetOrderByLocationAndSport(lat1,lng1 float64,sportType string)  {

}
//创建拼单
func CreateOrder(envelope model.Envelope)  {
//拼单成功后
}
func JoinOrder(){
	//在redis修改peoplenum，
	//在数据库中生成拼团者和拼团的联系

	//如果满了就拼团成功，并且拉群，websocket
}
