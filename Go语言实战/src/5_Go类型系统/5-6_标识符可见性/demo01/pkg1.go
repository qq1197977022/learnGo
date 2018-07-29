package demo01

import (
//"fmt"
//"5_Go类型系统/5-6_标识符可见性/pkg1"
)

func main() {
	//a := demo01.alertCount(10)	//访问pkg1包外不可见标识符
	//fmt.Println(a)

	/*
		# command-line-arguments
		src\5_Go类型系统\5-6_标识符可见性\demo01.go:10:2: cannot refer to unexported name demo01.alertCount
		src\5_Go类型系统\5-6_标识符可见性\demo01.go:10:2: undefined: demo01.alertCount
	*/

	//解决方案: 参见pkg2.go
}
