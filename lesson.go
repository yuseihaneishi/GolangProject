package main

import "fmt"

func main() {
	f := 1.11
	i := int(f)
	fmt.Printf("%T %v\n", i, i)

	s := []int{1, 2, 3}
	fmt.Println(s[1:])

	m := map[int]int{
		1: 200,
		2: 300,
	}
	fmt.Println(m)
}
