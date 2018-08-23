package day04

import (
	"fmt"
)

/*
处理恐慌
	1.内置函数panic, recover协助报告和处理运行时恐慌和程序定义的错误情况

	2.当执行函数F, 显式调用内置函数panic或运行时panic即终止函数F执行, F中deffer函数不受该运行时恐慌影响
	3.deffer执行完毕, 程序终止, 错误情况被报告, 包括panic的参数值
	4.上述终止序列称之为panicking(恐慌过程)

	5.内置函数recover允许程序管理处于恐慌中的goroutine, 其返回值是run-time panic信息和panic的参数值
*/
func HandlePanic() {
	squareRoot()
}

func squareRoot() { //求平方根
	defer handlePanic()

	for {
		fmt.Println("请输入一个的正数: ")
		sqrtNum := -1

		if fmt.Scan(&sqrtNum); sqrtNum >= 0 {
			fmt.Printf("value = %d, type = %T\n", sqrtNum, sqrtNum)
		} else {
			//panic("...你TM是不是傻?...")
			a := 0
			fmt.Println(3 / a)
		}
	}
}
func handlePanic() {
	if err := recover(); err != nil {
		fmt.Println("------", err, "------")
	} else {
		fmt.Println(err)
	}
}
