package main

import (
	"fmt"
	"sync"
	"time"
)

var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		total.Lock()
		time.Sleep(time.Second)
		total.value += i
		fmt.Println(total.value)
		total.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()

	fmt.Println(total.value)
}
