package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

	testCases := []string{"/view/test", "/edit/test", "/save/test"}
	for _, testCase := range testCases {
		fss := r2.FindStringSubmatch(testCase)
		if fss != nil {
			fmt.Println(fss, fss[0], fss[1], fss[2])
		} else {
			fmt.Printf("No match for input: %s\n", testCase)
		}
	}
}
