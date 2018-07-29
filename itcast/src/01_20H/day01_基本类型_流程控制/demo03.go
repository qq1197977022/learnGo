package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := 3
	p := &v
	fmt.Printf("%p\n", p)  //指针变量p, 保存变量v的地址
	fmt.Printf("%v\n", *p) //通过指针变量p, 访问变量v对应内存

	p1 := &p
	fmt.Printf("%p\n", p1)  //指针变量p1, 保存指针变量p的地址
	fmt.Printf("%p\n", *p1) //通过指针变量p1, 访问指针变量p对应内存
	fmt.Println(reflect.TypeOf(*p1))
	fmt.Println(reflect.TypeOf(*p1).Kind())
	fmt.Printf("%d\n", **p1) //通过指针变量*p1, 访问变量v对应内存
}
