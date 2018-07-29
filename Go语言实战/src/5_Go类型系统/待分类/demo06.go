package 待分类

import (
	"fmt"
)

func main() {
	var m1 = make(map[int]int, 5) //声明长度为5的空切片
	/*fmt.Println(m1, len(m1), m1 == nil)
	var m2 map[string]string	//声明nil映射
	fmt.Println(m2, len(m2), m2 == nil)*/

	for i := 0; i < 10; i++ {
		m1[i] = i + 1
	}
	//fmt.Println(m1)

	m1[0] = 0
	m1[1] = 0
	delete(m1, 0)
	delete(m1, 1)

	for key, value := range m1 {
		fmt.Printf("%v\t%v\n", key, value)
	}
}

/*
1.数组和映射的底层都是数组
*/
