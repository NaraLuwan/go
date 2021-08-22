package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["a"] = 1
	map1["b"] = 2
	fmt.Println(map1)

	map2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(map2)

	fmt.Println(map1["a"] == map2["b"]) // false
}
