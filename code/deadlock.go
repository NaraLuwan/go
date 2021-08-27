package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Start...")

	wg := sync.WaitGroup{}
	wg.Add(50)
	for i := 0; i < 10; i++ {
		i := i
		for j := 0; j < 5; j++ {
			j := j
			go func() {
				defer wg.Done()
				fmt.Printf("i: [%d], j: [%d] \n", i, j)
			}()
		}
	}
	wg.Wait()

	fmt.Println("Done!")
}
