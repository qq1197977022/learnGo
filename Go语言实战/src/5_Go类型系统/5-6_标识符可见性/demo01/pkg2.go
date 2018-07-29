package demo01

import (
	"5_Go类型系统/5-6_标识符可见性/demo01/pkg2"
	"fmt"
)

func main() {
	a := pkg2.New(10) //通过工厂函数,访问pkg2包内非导出类型
	fmt.Println(a)
}
