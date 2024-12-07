package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("lesson.go")
	defer file.Close()
	date := make([]byte, 100)
	file.Read(date)
	fmt.Println(string(date))
}
