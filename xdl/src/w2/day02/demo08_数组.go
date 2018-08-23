package day02

import "fmt"

func Array() {
	arr1 := [10]int{2: 333, 0: 222, 7: 100} //数组字面, 元素键值表示其下标
	fmt.Println(arr1)

	arr2 := [5]interface{}{"元组", "列表", 110, true, 3.14} //接口类型数组, 类似Python中的元组, 列表
	fmt.Println(arr2)

	fmt.Printf("数组型变量地址即其首元素地址: arr2Addr = %p\t arr2[0]Addr = %p", &arr2, &arr2[0])
	/*
		1.数组字面, 元素键值表示其下标
		2.接口类型数组, 类似Python中的元组, 列表
		3.对于类似数据结构, 学习点无非CURD
		4.数组型变量地址即其首元素地址
	*/
}
