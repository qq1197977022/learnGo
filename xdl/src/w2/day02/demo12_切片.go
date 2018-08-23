package day02

import "fmt"

/*
1.切片类型
	1.本质是其底层数组描述符
2.切片字面
3.切片表达式
	1.切片表达式从/字符串/数组/数组指针/切片构造一个子串/切片
	2.变形
		1.简单切片表达式
		2.完整切片表达式
*/
func Sli() {
	fmt.Println("------原切片栈存标头值中指针指向新底层数组------")
	m1()
	fmt.Println("------创建新切片------")
	m2()
}
func m2() {
	arr := [2]int{0, 1}
	sli := arr[:]
	fmt.Printf("arr:\t %p\t %p\n", &arr, &arr[0])
	fmt.Printf("sli:\t %p\t %p\n", &sli, &sli[0])
	sli1 := append(sli, 1)
	fmt.Printf("sli1:\t %p\t %p\n", &sli1, &sli1[0])
	sli2 := append(sli1, 2, 3, 4)
	fmt.Printf("sli2:\t %p\t %p\n", &sli2, &sli2[0])
}

func m1() {
	arr := [2]int{0, 1}
	sli := arr[:]

	fmt.Printf("arr:\t %p\t %p\n", &arr, &arr[0])
	fmt.Printf("sli:\t %p\t %p\n", &sli, &sli[0])

	sli = append(sli, 1, 22, 33, 44)
	fmt.Printf("sli:\t %p\t %p\n", &sli, &sli[0])

	sli = append(sli, 2, 3, 4)
	fmt.Printf("sli:\t %p\t %p\n", &sli, &sli[0])
}
