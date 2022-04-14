package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		timeoutWithBuffer(doBadThing)
	}
	time.Sleep(time.Second * 2)
	fmt.Printf("总协程数：%d", runtime.NumGoroutine())
}

func timeoutWithBuffer(f func(chan bool)) error {
	done := make(chan bool, 1)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func doBadThing(done chan bool) {
	time.Sleep(time.Second)
	// done <- true
}
