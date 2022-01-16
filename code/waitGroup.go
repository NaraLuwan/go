package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1000)
	var iList []int
	var iList2 []int
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		ii := i
		go func() {
			defer wg.Done()
			fmt.Println(ii)
			lock.Lock()
			defer lock.Unlock()
			iList = append(iList, ii)
			iList2 = append(iList2, ii)
		}()
	}
	wg.Wait()
	fmt.Println("-------------------------")
	fmt.Println(len(iList))
	fmt.Println(len(iList2))
}
