package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./lesson.go")
	if err != nil {
		log.Fatalln("Error")
	}
	defer file.Close()
	date := make([]byte, 100)
	count, err := file.Read(date)
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(date))

	if err = os.Chdir("test"); err != nil {
		log.Fatalln("Error")
	}
}
