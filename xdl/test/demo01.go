package main

import (
	"fmt"
	"sort"
)

func main() {
	sli := []int{1, 3, 2, 5, 4, 2}
	fmt.Println(sli)
	sort.Ints(sli)
	fmt.Println(sli)
}
