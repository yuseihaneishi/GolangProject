package main

import "fmt"

func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func main() {
	foo(1, 2, 3, 4, 5)
	foo(1, 2, 3)
}
