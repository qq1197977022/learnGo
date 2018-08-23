package 待分类

import "fmt"

func main() {
	var arr = [1e6]int{}

	foo1(arr) //值传递
	fmt.Println("+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+")
	foo2(&arr) //址传递 ~ &操作符获取地址
}

func foo1(arr [1e6]int) {
	fmt.Println(arr)
}

func foo2(arr *[1e6]int) {
	fmt.Println(arr)
}
