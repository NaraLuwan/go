package main

import (
    "fmt"
    "encoding/json"
)

type Person struct {
    Name string
    Age int
}

func main(){
    // test_struct()
    test_json()
}

func test_struct(){
    var p1 Person
    p1.Name = "Tom"
    p1.Age = 30
    fmt.Println("p1=", p1)

    var p2 = Person{Name: "Jack", Age: 18}
    fmt.Println("p2=", p2)

    p3 := Person{Name: "Luwan", Age: 20}
    fmt.Println("p3=", p3)

    p4 := struct{
        Name string
        Age int
    }{Name: "匿名", Age: 50}
    fmt.Println("p4=", p4)
}

// 生成json
type Result struct{
    Code int `json:"code"`
    Message string `json:"msg"`
}

func test_json(){
    var res Result
    res.Code = 200
    res.Message = "ok!"

    // 序列化
    jsons, errs := json.Marshal(res)
    if errs != nil {
        fmt.Println("json marshal err! errMsg:", errs)
    }
    fmt.Println("json data:", string(jsons))
}

