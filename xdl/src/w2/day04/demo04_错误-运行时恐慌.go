package day04

var arr = [3]int{0, 1, 2}

func Error_RunTimePanic() {
	err()
	runTimePanic()
}
func err() {
	/*
		Error接口是一个用于表示错误情况的惯例接口, nil表示没有错误
	*/
}
func runTimePanic() {
	/*
		执行类似数组越界的错误会触发一个等价于实现了接口类型runtime.Error的值的内建函数panic调用的运行时恐慌
		该类型满足预声明接口类型error
		表示不同运行时错误情形的准确错误值不确定
	*/
}
