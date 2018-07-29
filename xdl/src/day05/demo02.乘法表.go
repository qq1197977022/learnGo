package day05

import "fmt"

func MultTable() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

	fmt.Println("\n------------------------------------")
	for i := 10; i > 0; i-- {
		for j := 9; j >= i; j-- {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
