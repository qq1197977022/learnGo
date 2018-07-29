package main

import "fmt"

func main() {
	var i interface{} = 3
	fmt.Printf("i = %v\t%T\n", i, i)
	i = 3.0
	fmt.Printf("i = %v\t%T\n", i, i)
	i = '3'
	fmt.Printf("i = %v\t%T\n", i, i)
	i = "3"
	fmt.Printf("i = %v\t%T\n", i, i)
	i = true
	fmt.Printf("i = %v\t%T\n", i, i)

	t()
	t(1, 'A', "B")

	i = 3
	v, ok := i.(int) //Type assertion
	fmt.Println(v, ok)

	switch i.(type) { //Type switch
	case int:
		fmt.Println("i is int type variable")
	}

}
func t(args ...interface{}) {
	fmt.Printf("%v\t%T\n", args, args)
}

/*
面向对象编程 ~ 实际上Go语言没有此概念
	2.继承 ~ 嵌入字段
	1.封装 ~ 方法
	3.多态 ~ 接口
*/

/*
1.空接口没有方法, 因此所有类型都实现了空接口, 即空接口相当于是一个万能类型
*/
