package main

import("fmt")

func main(){
    // 常量声明
    const name string = "Tom"
    fmt.Println(name)
    const age int = 30
    fmt.Println(age)

    const name_1, name_2 string = "Tom", "Jay"
    fmt.Println(name_1, name_2)

    const name_3, age_1 = "Tom", 30
    fmt.Println(name_3, age_1)

    // 变量声明
    var count_1 uint8 = 31
    var count_2 = 32
    count_3 := 33
    fmt.Println(count_1, count_2, count_3)

    var count_4, count_5, count_6 = 34, 35, 36
    fmt.Println(count_4, count_5, count_6)

    count_6 = 37
    fmt.Println(count_6)

    // 输出方法
    fmt.Print("输出到控制台不换行")
    fmt.Println("---")
    fmt.Println("输出到控制台并换行")
    fmt.Printf("name=%s, age=%d\n", "Tom", 30)
    fmt.Printf("name=%s, age=%d, height=%v\n", "Tom", 30, fmt.Sprintf("%.2f", 180.567))
}