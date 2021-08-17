package main

import (
	"fmt"
)

func main()  {

	var array [10]int
	var slice = array[5:6]
	fmt.Println("lenth of slice: ", len(slice))
	fmt.Println("capacity of slice: ", cap(slice))
	fmt.Println(&slice[0] == &array[5])
//	b:= slice[1:3]
	//fmt.Println(&slice[1] == &b[0])
	orderLen:=5
	order:=make([]int,2*orderLen)

	pollorder:=order[:orderLen:orderLen]
	lockorder:=order[orderLen:][:orderLen:orderLen]
	fmt.Println(len(lockorder))
	fmt.Println(len(pollorder))

}
