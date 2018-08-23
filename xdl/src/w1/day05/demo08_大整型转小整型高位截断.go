package day05

import "fmt"

func BigToSmallInt() {
	var big int64 = 1234567890

	fmt.Printf("big: %b\t %d\n", big, big)
	fmt.Printf("BigToSmallInt: %d\t %b\n", uint8(big), uint8(big))

	/*
		十进制				二进制
		1234567890			1001001100101100000001011010010
		210					11010010

		1001001100101100000001011010010 ---> 高位截断 ---> 11010010 ---> 无符号十进制:2+16+64+128=210
	*/
}
