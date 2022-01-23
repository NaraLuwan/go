package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

//type resp struct {
//	sync.RWMutex
//	Count int
//}
//
//func main() {
//	resp := &resp{
//		RWMutex: sync.RWMutex{},
//		Count:     0,
//	}
//	wg := sync.WaitGroup{}
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		ii := i
//		go func() {
//			wg.Done()
//			time.Sleep(time.Second)
//			resp.Lock()
//			fmt.Printf("第%d次处理\n", ii)
//			resp.Count += 1
//			resp.Unlock()
//		}()
//	}
//	wg.Wait()
//	fmt.Printf("main end, resp: %+v", resp)
//}

func main() {
	now := time.Now()
	wg := &sync.WaitGroup{}
	lock := &sync.Mutex{}
	var res []string
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			if i == 50 {
				time.Sleep(time.Second)
			}
			time.Sleep(time.Second)
			lock.Lock()
			res = append(res, strconv.FormatInt(int64(i), 10))
			lock.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Since(now))
}

//func parse(i int, ch chan string, wg *sync.WaitGroup) {
//	defer wg.Done()
//	ch <- strconv.FormatInt(int64(i), 10)
//	time.Sleep(time.Second)
//}

//func main() {
//	now := time.Now()
//	arr := []int{1, 2, 3, 4, 5}
//	wg := &sync.WaitGroup{}
//	ch := make(chan string, len(arr))
//	var res []string
//	for i := range arr {
//		wg.Add(1)
//		go parse(i, ch, wg)
//	}
//	wg.Wait()
//	close(ch)
//	for s := range ch {
//		res = append(res, s)
//	}
//	fmt.Println(time.Since(now))
//}
