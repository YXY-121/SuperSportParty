package service

import (
	"apiproject/SuperSport/model"
	"apiproject/SuperSport/repository"
	"apiproject/common"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"time"
)

//
type OrderService struct {
	orderRepository  *repository.OrderRepository
}


//根据位置获取拼团活动
func (o*OrderService)GetOrderByLocation (lat1,lng1 float64)  {

}
//根据位置、运动类型获取拼团活动
func (o*OrderService)GetOrderByLocationAndSport(lat1,lng1 float64,sportType string)  {

}
//创建拼单
func (o*OrderService)CreateOrder(createrId string,orderLocation string,sportType string,description string,peopleNumber int,endTime time.Time)  {

//拼单成功后
	orderId:=uuid.UUID{}.String()
	order:=model.Order{
		CreaterId: createrId,
		OrderLocation:orderLocation,
		OrderId:orderId,
		SportType: sportType,
		Description: description,
		PeopleNumber: peopleNumber,
		EndTime: endTime,
		CreateTime: time.Now(),
	}
	//写入数据库
	o.orderRepository.AddOrder(order)
	//并将order信息存到redis里去,防止死锁
	reply,_:=common.RedisDB.Do("set",orderId,1,"ex",10,"nx")
	if(reply==nil){
		fmt.Errorf("加锁失败")
	}else{
		fmt.Println("成功加锁")
	}
}
func (o*OrderService)JoinOrder(userId string,orderId string){
	//在redis修改peoplenum，
	count:=0//count是试图加锁的次数，设置count为5时，就睡一会
	for{
		//todo 这里后面要用lua来优化，防止把别人的锁给删了
		reply,_:=common.RedisDB.Do("set",orderId,userId,"ex",10,"nx")
		if(reply==nil){
			count++
			if count==5{
				time.Sleep(time.Second*2)
				count=0
			}
			//抢到锁了，扣减库存,扣完后归还锁
		}else if reply=="OK"{
			//扣减库存
			o.orderRepository.DelOrderPeopleNumber(orderId,1);

			//生成记录
			NewUserOrder().UserOrderRepository.CreateUserOrderRecord(userId,orderId)

			//解锁
			common.RedisDB.Do("del",orderId,userId,"ex",10,"nx")

			break;
		}

	}


	//如果满了就拼团成功，并且拉群，websocket
}
