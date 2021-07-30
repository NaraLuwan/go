package main

import "fmt"

type A struct {
	MyInt int
}

func (a A) method() {
	println("this is A struct method")
}

type A2 A

func main() {
	var a A = A{
		MyInt: 10,
	}
	a.method()

	var a2 A2 = A2{
		MyInt: 20,
	}
	fmt.Println(a2.MyInt)
	// 编译失败，自定义新类型没有基础类型A的方法
	//a2.method()
}
