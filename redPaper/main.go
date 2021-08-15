package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math/rand"
	"time"
)
type Reward struct {
	Count          int   //个数
	Money          int   //总金额(分)
	RemainCount    int   //剩余个数
	RemainMoney    decimal.Decimal   //剩余金额(分)
	BestMoney int   //手气最佳金额
	BestMoneyIndex int   //手气最佳序号
	MoneyList      []int //拆分列表
}



var multiple = decimal.NewFromFloat(100.0)


func DoubleAverage(count, amount int64) int64 {
	//最小钱
	min := int64(1)

	if count == 1 {
		//返回剩余金额
		return amount
	}

	//计算最大可用金额,min最小是1分钱,减去的min,下面会加上,避免出现0分钱
	max := amount - min*count
	//计算最大可用平均值
	avg := max / count
	if avg<0 {
		avg=0
	}
	//二倍均值基础加上最小金额,防止0出现,作为上限
	avg2 := 2*avg + min
	//随机红包金额序列元素,把二倍均值作为随机的最大数
	rand.Seed(time.Now().UnixNano())
	//加min是为了避免出现0值,上面也减去了min
	x := rand.Int63n(avg2) + min

	return x
}


func main() {
	//初始10个红包, 10000分 = 100元钱
	count, amount := int64(10), int64(5)
	//剩余金额
	remain := amount
	//验证红包算法的总金额,最后sum应该==amount
	sum := int64(0)
	//进行发红包
	for i := int64(0); i < count; i++ {
		x := DoubleAverage(count-i, remain)
		//金额减去
		remain -= x
		//发了多少钱
		sum += x
		//金额转成元
		fmt.Println(i+1, "=", float64(x)/float64(100))
	}
	fmt.Println()
	fmt.Println("总和 ", sum)
}

// 随机红包
// remainCount: 剩余红包数
// remainMoney: 剩余红包金额（单位：分)
func randomMoney(remainCount int , remainMoney decimal.Decimal )decimal.Decimal{
	if remainCount == 1{
		fmt.Println("最后一个",remainMoney)

		return remainMoney
	}

//	rand.Seed(time.Now().UnixNano())
//	var min = 0.0
	decimalcount:=decimal.NewFromInt(int64(remainCount))
	max := remainMoney .Div(decimalcount).Mul(decimal.NewFromInt(int64(2)))
	fmt.Println("max", max)
	fmt.Println(max.IntPart())

	r:=rand.Intn(int(max.IntPart()))

	decimalR:=decimal.NewFromInt(int64(r)).Add(decimal.NewFromInt(int64(1)))
	fmt.Println("r",decimalR)

//	money :=decimalR.Mul(max)
		//r*float64(max)+min
	//fmt.Println("钱 ", r)


	//rand.Intn(max) + min
	return decimalR
}

// 发红包
// count: 红包数量
// money: 红包金额（单位：分)
func redPackage(count int, money decimal.Decimal)  {
	var count1 decimal.Decimal

	for i := 0; i < count; i++ {
		m := randomMoney(count - i, money)
		count1=count1.Add(m)
		money=money.Sub(m)

	}
	fmt.Println("总共花了",count1)
}