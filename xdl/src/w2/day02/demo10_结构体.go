package day02

import "fmt"

type human struct {
	name   string
	age    int
	gender bool
}

func StructType() {
	tom1 := human{}
	var tom2 human
	tomPtr := new(human) //为human类型变量分配内存, 并初始化, 返回*S类型指针指向分配的内存

	fmt.Println(tom1)
	fmt.Println(tom2)
	fmt.Printf("%T\t %v\t %#v", tomPtr, tomPtr, tomPtr)
}
