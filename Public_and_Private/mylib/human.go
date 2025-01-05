package mylib

import "fmt"

var Public string = "Public"

// 変数名や関数名etcをキャピタルにしないとパッケージ内でしか使用することができない
var Private string = "private"

type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Human!")
}
