package main

import "fmt"

func main() {
	// 一维数组
	var arr_1 [5]int
	fmt.Println(arr_1)

	var arr_2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_2)

	arr_3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_3)

	arr_4 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr_4)

	arr_5 := [5]int{0: 3, 1: 5, 4: 6}
	fmt.Println(arr_5)

	// 二维数组
	var arr_6 = [3][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_6)

	arr_7 := [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_7)

	arr_8 := [...][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {0: 3, 1: 5, 4: 6}}
	fmt.Println(arr_8)

	// 长度不可变，运行会报错：invalid array index 4 (out of bounds for 3-element array)
	var arr_9 = [3]int{1, 2, 3}
	// arr_9[4] = 4    // 去掉这行就可以
	fmt.Println(arr_9)

	// 数组是值类型问题，在函数中传递的时候是传递的值，如果传递数组很大，这对内存是很大开销
	// 运行输出[1, 2, 3, 4, 5]
	var arr_10 = [5]int{1, 2, 3, 4, 5}
	modifyArr(arr_10)
	fmt.Println(arr_10)

	// 运行输出[1, 20, 3, 4, 5]
	var arr_11 = [5]int{1, 2, 3, 4, 5}
	modifyArr2(&arr_11)
	fmt.Println(arr_11)

	// 数组赋值问题，同样类型的数组（长度一样且每个元素类型也一样）才可以相互赋值，反之不可以
	// 运行会报错：cannot use arr_12 (type [5]int) as type [6]int in assignment
	var arr_12 = [5]int{1, 2, 3, 4, 5}
	var arr_13 [5]int = arr_12
	// var arr_14 [6] int = arr_12    // 去掉这行就可以
	fmt.Println(arr_13)
	// fmt.Println(arr_14)    // 去掉这行就可以
}

func modifyArr(arr [5]int) {
	arr[1] = 20
}

func modifyArr2(arr *[5]int) {
	arr[1] = 20
}
