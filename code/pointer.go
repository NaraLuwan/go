package main

import (
	"fmt"
)

func testString(s string) {
	fmt.Printf("inner: %v, %v\n", s, &s)
	s = "b"
	fmt.Printf("inner: %v, %v\n", s, &s)
}

func testMap(m map[string]string) {
	fmt.Printf("inner: %v, %p\n", m, m)
	m["a"] = "11"
	fmt.Printf("inner: %v, %p\n", m, m)
}

func main() {
	s := "a"
	fmt.Printf("outer: %v, %v\n", s, &s)
	testString(s)
	fmt.Printf("outer: %v, %v\n", s, &s)

	m := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	fmt.Printf("outer: %v, %p\n", m, m)
	testMap(m)
	fmt.Printf("outer: %v, %p\n", m, m)
}
