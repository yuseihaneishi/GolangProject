package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(m)
	fmt.Println(m["a"])
	m["banana"] = 300
	fmt.Println(m)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)
}
