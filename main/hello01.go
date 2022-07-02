package main

import "fmt"

func main2(){
	//声明一个数组
	var balance [5] float32
	var balance2 = [5]float32 {1,2,3,4,5}
	var balance3 = []float32 {1,2,3,4,5}

	balance[0] = 10

	fmt.Println(balance)
	fmt.Println(balance2)
	fmt.Println(balance3)

	// 指针
	var a int = 20
	var ip *int = &a

	//地址
	fmt.Println(&a)
	//地址
	fmt.Println(ip)
	//使用指针访问值
	fmt.Println(*ip)
}