package main

import (
	"fmt"

	"github.com/hb-0/testgo1/pack1"
)

func main() {
	// 定义一个结构体
	var myStruct1 pack1.MyStruct
	myStruct1.Age = 18
	myStruct1.Name = "张三"
	fmt.Println(myStruct1)

	// // 定义一个结构体
	// var myStruct2 pack2.myStruct
	// myStruct2.Age = 18
	fmt.Println("Hello World2!")
}
