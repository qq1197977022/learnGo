package day05

import (
	"fmt"
	"unsafe"
)

func MeasureSize() { //用到unsafe包
	var n1 int
	n2 := 3
	var byt byte = 3 //uint8别名
	var n3 int8 = 3

	var c1 rune = 'A'
	c2 := 'B'

	var b1 bool = true
	b2 := false

	var str1 = "A"
	str2 := "AB"

	var arr1 = [3]int{1, 2, 3}
	arr2 := [...]string{"我", "狂草"}
	arr3 := [...]int{}
	arr4 := [3]int{}

	fmt.Printf("%d\t %T\n", unsafe.Sizeof(arr1), arr1)
	fmt.Printf("%d\t %T\n", unsafe.Sizeof(arr2), arr2)
	fmt.Printf("%d\t %T\n", unsafe.Sizeof(arr3), arr3)
	fmt.Printf("%d\t %T\n", unsafe.Sizeof(arr4), arr4)

	fmt.Println("------")
	fmt.Printf("%d\t %T\n", unsafe.Sizeof(str1), str1)
	fmt.Printf("%d\t %T\n", unsafe.Sizeof(str2), str2)

	fmt.Println("------")
	fmt.Printf("%d\t %T\n", unsafe.Sizeof(b1), b1)
	fmt.Printf("%d\t %T\n", unsafe.Sizeof(b2), b2)

	fmt.Println("------")
	fmt.Printf("%d\t %T\n", unsafe.Sizeof(MeasureSize), MeasureSize)

	fmt.Println("------")
	fmt.Printf("%d\t %T\n", unsafe.Sizeof(c1), c1)
	fmt.Printf("%d\t %T\n", unsafe.Sizeof(c2), c2)

	fmt.Println("------")
	fmt.Printf("%d\t %T\n", unsafe.Sizeof(byt), byt)
	fmt.Printf("%d\t %T\n", unsafe.Sizeof(n1), n1)
	fmt.Printf("%d\t %T\n", unsafe.Sizeof(n2), n2)
	fmt.Printf("%d\t %T\n", unsafe.Sizeof(n3), n3)
}
