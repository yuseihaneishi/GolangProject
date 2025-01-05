package main

import (
	"awesomeProject/package/package/mylib"
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))
}
