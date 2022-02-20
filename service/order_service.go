package service

import (
	"apiproject/common"
	"apiproject/model"
	"apiproject/repository"
	"context"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"

	"github.com/gomodule/redigo/redis"

	"time"
)

//
type OrderService struct {
	orderRepository *repository.OrderRepository
}

func NewOrderService() *OrderService {
	return &OrderService{
		repository.NewOrderRepository(),
	}
}

//根据位置获取拼团活动
//根据区域划分出
func (o *OrderService) GetOrderByLocation(location string) []model.Order {
	return o.orderRepository.GetOrdersByLocation(location)
}

//根据位置、运动类型获取拼团活动
func (o *OrderService) GetOrderByLocationAndSport(lat1, lng1 float64, sportType string) {

}

// 生成区间[-m, n]的安全随机数
func RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max!")
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}

//创建拼单
func (o *OrderService) CreateOrder(order model.Order) {

	//拼单成功后
	//写入数据库
	o.orderRepository.CreateOrder(order)
	//	biteOrder, _ := json.Marshal(order)

	//并将order信息存到redis里去,防止死锁
	//	fmt.Println("存redis了嘛")
	//	common.RedisDB.Do("set", order.OrderId+"Item", biteOrder, "EX", RangeRand(60, 1000))

}

const unlock = `
if redis.call("get",KEYS[1])==ARGV[1]
then
	redis.call("DEL",KEYS[1])
	return 1
else
	return 0
end`

//一定要写end 不然就不会实现T.T
const test = `
local a= tostring(ARGV[1])
if redis.call("get",KEYS[1])==a
then
return ARGV[1]
else
return ARGV[1]
end
`
const lock = `return	redis.call("set",KEYS[1],ARGV[1],"EX",5,"NX")`

func (o *OrderService) GrabOrder(userId string, orderId string) {
	fmt.Println(orderId)
	//在redis修改peoplenum，
	count := 0 //count是试图加锁的次数，设置count为5时，就睡一会
	for {

		lockLua := redis.NewScript(1, lock)
		unlockLua := redis.NewScript(1, unlock)
		//expireTimeLua:=redis.NewScript(1,expireTime)

		reply, _ := lockLua.Do(common.RedisDB, orderId, userId)

		if reply == nil {
			count++
			time.Sleep(time.Second * 2)

			if count == 5 {
				fmt.Println("抢锁失败")
				break
			}
			//抢到锁了，扣减库存,扣完后归还锁
		} else if reply == "OK" {
			//开启协程延时开门狗
			ctx, cancel := context.WithCancel(context.Background())
			go o.WatchDog(orderId, userId, ctx)

			//扣减库存
			o.orderRepository.DelOrderPeopleNumber(orderId, 1)

			//生成记录
			NewUserOrder().UserOrderRepository.CreateUserOrderRecord(userId, orderId)
			//模拟超长的业务时间
			//time.Sleep(20*time.Second)
			//解锁
			unlockReply, _ := redis.Int(unlockLua.Do(common.RedisDB, orderId, userId))
			cancel()
			if unlockReply == 1 {
				fmt.Println(userId, "解锁")
			}

			break
		}

	}

	//如果满了就拼团成功，并且拉群，websocket
}

const expireTime = `
if redis.call("GET",KEYS[1])==ARGV[1]
then
redis.call("EXPIRE",KEYS[1],30)
return 1
else 
return 0
end`

func (o *OrderService) WatchDog(key string, value string, ctx context.Context) {
	expireTimeLua := redis.NewScript(1, expireTime)
	for {
		select {
		case <-ctx.Done():
			fmt.Println(value, "看门狗要走辣")
			return
		default:
			reply, _ := redis.Int(expireTimeLua.Do(common.RedisDB, key, value))
			if reply == 1 {
				fmt.Println(value, "加时成功")
			}
			time.Sleep(10 * time.Second)

		}
	}

}

func (o *OrderService) GetOrderById(orderId string) model.Order {
	return o.orderRepository.GetOrderByOrderId(orderId)
}

func (o *OrderService) GetOrderByTypeAndLocation(orderType string, location string) []model.Order {
	return o.orderRepository.GetOrderByTypeAndLocation(orderType, location)
}
