package main

import "fmt"

func main() {
	a := 10
	fmt.Println(&a)
	*&a = 100
	fmt.Println(a)

	var b *int
	b = new(int)
	*b = 10
	fmt.Println(*b)

	var c map[string]int
	c = make(map[string]int, 10)
	c["k1"] = 100
	fmt.Println(c)

	var d = make(map[string]int, 10)
	d["k1"] = 100
	d["k2"] = 200
	d["k2"] = 300
	fmt.Println(d)
	fmt.Println(len(d))
}
