package main

import (
	"fmt"
	"time"
)

func main() {
	TIME_ZERO := "2006-01-02 15:04:05"
	// 时间戳转时间字符串 (int64 —>  string)
	timeUnix := time.Now().Unix() //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format(TIME_ZERO)
	fmt.Println(formatTimeStr) //打印结果：2017-04-11 13:30:39
}
