package main

import (
	"fmt"
	"sort"
)

func main() {
	test := []int{3, 5, 8, 2, 4, 1, 9, 4, 6, 3}
	fmt.Println(test)
	sort.Ints(test)
	fmt.Println(test)
}
