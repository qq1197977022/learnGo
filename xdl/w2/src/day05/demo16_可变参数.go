package day05

import "fmt"

func Variadic() {
	sli1 := []int{1, 2, 3}
	demo01()
	demo01(1, 2)
	demo01(sli1...)
	fmt.Printf("%v, %T %p\n", sli1, sli1, sli1)
	/*
	   1.参数
	   	1.实参: arguments
	   	2.形参: parameters
	   1.可变参数
	   	1.特指形参parameters
	   	2.在函数内相当于[]T
	   	3.可变形参实参
	   		1.默认是nil
	   		2.[]T, 实参arguments存储在对应底层数组中
	   	4.若实参arguments是[]T, 则直接作为可变形参parameters的值, 若实参arguments后跟..., 则不会创建新的切片, 实参arguments和形参parameters共享同一底层数组
	*/
}

func demo01(variadic ...int) {
	fmt.Printf("%v, %T, %p\n", variadic, variadic, variadic)
}
