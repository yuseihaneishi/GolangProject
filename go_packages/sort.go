package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{5, 3, 2, 8, 7}
	s := []string{"z", "d", "a", "h"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 39},
		{"Yusei", 20},
		{"Nami", 25},
		{"Minami", 18},
	}
	fmt.Println(i, s, p)
	sort.Ints(i)
	sort.Strings(s)
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	fmt.Println(i, s, p)
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(i, s, p)
}
