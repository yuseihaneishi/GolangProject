package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b []string = []string{"hello", "world"}
	b = append(b, "halloe")
	fmt.Println(b[0][1])
}
