package main

import "fmt"

func add(x, y int32) int32 {
	if x < 0 || y < 0 {
		return 0
	}
	return x + y
}

func testAdd(x, y int32, op func(int32, int32) int32) (int32, string) {
	return op(x, y), ""
}

func main() {
	// 函数作为参数
	fmt.Println(testAdd(10, 10, add))

	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20)

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(1, 2)
}
