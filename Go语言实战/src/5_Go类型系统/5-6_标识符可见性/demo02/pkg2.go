package main

import (
	"5_Go类型系统/5-6_标识符可见性/demo02/pkg2"
	"fmt"
)

func main() {
	a := pkg2.Admin{
		Rights: 10,
	}
	a.Name = "杨一帆"
	a.Email = "1192@qq.com"
	fmt.Println(a)
}
