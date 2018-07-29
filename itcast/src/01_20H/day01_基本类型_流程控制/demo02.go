package main

import (
	"fmt"
	"os"
)

func main() {
	f1()
	fmt.Println("----------------------------")
	f2()

	fmt.Println(os.Args, len(os.Args))
}
func f1() {
	x, y := 3, 4
	defer func() {
		fmt.Printf("x = %d, y = %d\n", x, y)
	}()

	x, y = 33, 44
	fmt.Printf("x = %d, y = %d\n", x, y)
}

func f2() {
	x, y := 3, 4
	defer func(x, y int) {
		fmt.Printf("x = %d, y = %d\n", x, y)
	}(x, y)

	x, y = 33, 44
	fmt.Printf("x = %d, y = %d\n", x, y)
}
