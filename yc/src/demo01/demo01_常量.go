package demo01

import "fmt"

func Demo01() {
	const (
		a = iota
		b = iota
		c
		d
		m = 3
		p //p, i = 3 常量声明语法决定, 参见Demo02
		i
		q = iota //q = 7 预定义标识符iota特性决定
		h
	)
	fmt.Println(p, c, q, h)
	fmt.Println(i, m)
}
func Demo02() {
	/*
		1.常量声明, 后续省略, 默认等于上一常量值
		2.上述iota也是常量 ~  the predeclared identifier iota represents successive untyped integer constants
	*/
	const (
		a = 3
		b
		c
	)
	fmt.Println(a, b, c)
}
