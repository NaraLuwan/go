package main

import (
	"fmt"
	"time"
)

func main() {

	go doTask()

	for {
		time.Sleep(time.Second * 1)
	}
}

func doTask() {
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("do task...")
		}
	}
}
