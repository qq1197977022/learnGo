package main

import (
	"5_Go类型系统/5-6_标识符可见性/demo02/pkg1"
	"fmt"
)

func main() {
	u := pkg1.User{
		Name:  "杨一帆",
		email: "1192@qq.com",
	}

	fmt.Println(u)
	/*
		# command-line-arguments
		src\5_Go类型系统\5-6_标识符可见性\demo02\pkg1.go:11:8: unknown field 'email' in struct literal of type pkg1.User (but does have pkg1.email)
	*/
}
