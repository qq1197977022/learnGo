package day05

import "fmt"

func MultTable() {
	for i := 1; i < 10; i++ { //行
		for j := 1; j <= i; j++ { //列
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

	fmt.Println("\n------------------------------------")
	for i := 10; i > 0; i-- { //行
		for j := 9; j >= i; j-- { //列
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
