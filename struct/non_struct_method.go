package main

import "fmt"

type Myint int

func (i Myint) Double() int {
	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", 1, 1)
	return int(i * 2)
}

func main() {
	myInt := Myint(10)
	fmt.Println(myInt.Double())
}
