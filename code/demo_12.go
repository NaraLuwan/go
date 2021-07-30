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
}
