package main

import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {
	// 精确到秒
	c := cron.New(cron.WithSeconds())

	//每秒一次
	spec := "*/1 * * * * ?"
	_, _ = c.AddFunc(spec, func() {
		fmt.Println("11111")
	})

	c.Start()
	//阻塞主线程停止
	select {}
}
