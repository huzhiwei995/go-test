package main

import "fmt"

func main() {
	//一维数组
	arrayOne := [4]int{1,2,3,4}
	arrayTwo := [4][2]int{{1,2},{3,4},{5,6},{7,8}}

	arrayOneThree := [4][2]int{1: {20, 21}, 3: {40, 41}}

	// 声明并初始化外层数组和内层数组的单个元素
	array := [4][2]int{1: {0: 20}, 3: {1: 41}}

	fmt.Println(arrayOne)
	fmt.Println(arrayTwo)
	fmt.Println(arrayOneThree)
	fmt.Println(array)
}