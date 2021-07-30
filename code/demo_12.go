package main

import (
	"encoding/json"
	"fmt"
)

/**
变量首字母大写才可以导出，才可以序列化
*/
type Stu struct {
	// `json:"name"` json key为name，如果没加这个标签则为Name
	Name string `json:"name"`
	Age  int
	High bool
	sex  string
	// 指针变量编码时自动转换为它所指向的值
	Class *Class
}

type Class struct {
	Name  string
	Grade int
}

/**
上面的成员变量都是已知的类型，只能接收指定的类型，比如string类型的Name只能赋值string类型的数据
但有时为了通用性，或使代码简洁，我们希望有一种类型可以接受各种类型的数据，并进行json编码。这就用到了interface{}类型
*/
type Stu2 struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	HIgh  interface{}
	sex   interface{}
	Class interface{} `json:"class"`
}

func main() {
	// 实例化一个结构
	stu := Stu{
		Name: "Tom",
		Age:  18,
		High: true,
		sex:  "男",
	}

	/**cla2 := Class{
		Name: "2班",
		Grade: 2,
	}*/

	cla := new(Class)
	cla.Name = "1班"
	cla.Grade = 3
	stu.Class = cla

	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("序列化错误")
	}

	fmt.Println(string(jsonStu))

	// 实例化一个结构
	stu2 := Stu{
		Name: "Tom",
		Age:  18,
		High: true,
		sex:  "男",
	}
	stu2.Class = cla

	jsonStu2, err := json.Marshal(stu2)
	if err != nil {
		fmt.Println("序列化错误")
	}

	fmt.Println(string(jsonStu2))
}
