package main

import "fmt"

func main(){
    var sli_1 [] int    // nil 切片
    fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_1),cap(sli_1),sli_1)
}