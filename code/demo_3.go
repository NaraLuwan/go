package main

import "fmt"

func main(){
    test_init()
    test_split()
    test_add()
    test_del()
}

// 切片定义
func test_init(){
    var sli_1 [] int    // nil 切片
    fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_1),cap(sli_1),sli_1)

    var sli_2 = [] int {}   // 空切片
    fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_2), cap(sli_2), sli_2)

    var sli_3 = [] int {1, 2, 3, 4, 5}
    fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_3),cap(sli_3),sli_3)

    sli_4 := [] int {1, 2, 3, 4, 5}
    fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_4),cap(sli_4),sli_4)

    var sli_5 [] int = make([] int, 5, 8)
    fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_5),cap(sli_5),sli_5)

    sli_6 := make([] int, 5, 9)
    fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_6),cap(sli_6),sli_6)
}

// 截取切片
func test_split(){
sli_7 := [] int {1, 2, 3, 4, 5, 6}
    fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_7),cap(sli_7),sli_7)

    fmt.Println("sli[1] ==", sli_7[1])
    fmt.Println("sli[:] ==", sli_7[:])
    fmt.Println("sli[1:] ==", sli_7[1:])
    fmt.Println("sli[:4] ==", sli_7[:4])

    fmt.Println("sli[0:3] ==", sli_7[0:3])
    fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_7[0:3]),cap(sli_7[0:3]),sli_7[0:3])

    // 长度为3 容量为4
    fmt.Println("sli[0:3:4] ==", sli_7[0:3:4])
    fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_7[0:3:4]),cap(sli_7[0:3:4]),sli_7[0:3:4])
}

// 追加切片
func test_add() {
	sli := [] int {4, 5, 6}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	sli = append(sli, 7)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	sli = append(sli, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	sli = append(sli, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

    // 自动扩容
	sli = append(sli, 10)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)
}

// 删除切片
func test_del(){
    sli := [] int {1, 2, 3, 4, 5, 6, 7, 8}
    fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

    //删除尾部 2 个元素
    fmt.Printf("len=%d cap=%d slice=%v\n",len(sli[:len(sli)-2]),cap(sli[:len(sli)-2]),sli[:len(sli)-2])

    //删除开头 2 个元素
    fmt.Printf("len=%d cap=%d slice=%v\n",len(sli[2:]),cap(sli[2:]),sli[2:])

    //删除中间 2 个元素
    sli = append(sli[:3], sli[3+2:]...)
    fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)
}