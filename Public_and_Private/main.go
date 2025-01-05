package main

import (
	"awesomeProject/mylib"
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	mylib.Say()
	person := mylib.Person{"Mike", 22}
	fmt.Println(person)

	fmt.Println(mylib.Public)
	//fmt.Println(mylib.pribate)
}
