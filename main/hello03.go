package main

import "fmt"

// 切片就是动态数组
// 动态数组有一些语法糖，方便处理数组

func main4(){

	var nums = []int {1, 2, 3}
	var nums2 = make([]int, 3, 5)

	fmt.Println(nums)
	fmt.Println(nums2)
	fmt.Println(nums[:2])

	var total = 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total: ", total)

	total = 0
	for index, num := range nums {
		total += num
		fmt.Println("index: ", index, "num: ", num)
	}

	fmt.Println("total: ", total)

}