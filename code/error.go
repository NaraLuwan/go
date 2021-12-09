package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	// errors.New()
	err1 := errors.New("new error!")
	fmt.Printf("errorType:%T, err:%v \n", err1, err1)

	// fmt.Errorf
	err2 := fmt.Errorf("error format error! msg:%s", "err msg")
	fmt.Printf("errorType:%T, err:%v \n", err2, err2)

	// fmt.Errorf && warpError
	err3 := fmt.Errorf("warp error format error! msg:%s, origin err:%w", "new msg", err2)
	fmt.Printf("errorType:%T, err:%v \n", err3, err3)

	// Unwrap err
	err22 := errors.Unwrap(err3)
	fmt.Println(err2 == err22) // true
	fmt.Printf("errorType:%T, err:%v \n", err22, err22)
}
