package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:

		fmt.Println(v * 2)
	case string:

		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", i)
	}
}

func main() {
	do(10)
	do("Mike")
	do(true)

	var i int = 10
	ii := float64(10)
	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", ii, ii)

}
