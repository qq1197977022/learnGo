package 待分类

import "fmt"

func main() {
	/*//sl:= make([]int, 5, 3)
	//fmt.Println(sl)

	//nil切片 ~ 没有对应底层数组
	var sl1 []int	//声明切片,指针nil,长度0,容量0
	fmt.Println(sl1, sl1 == nil, len(sl1), cap(sl1))

	//空切片 ~ 有对应的底层数组,但没有分配任何空间
	sl2 := make([]int, 0)	//调用内置函数make创建空切片,返回T而不是*T
	sl3 := []int{}	//空切片字面
	fmt.Println(sl2, sl2 == nil, len(sl2), cap(sl2))
	fmt.Println(sl3, sl3 == nil, len(sl3), cap(sl3))

	//向切片中追加元素
	s1 := append(sl1, 3)
	s2 := append(sl2, 3)
	fmt.Println(s1, len(s1), cap(s1), "\t", sl1, len(sl1), cap(sl1))
	fmt.Println(s2, len(s2), cap(s2), "\t", sl2, len(sl2), cap(sl2))*/

	//sl1 := []int{1, 2, 3, 4}
	//fmt.Println(sl1, len(sl1), cap(sl1))
	//
	//sl2 := append(sl1, 5)
	//fmt.Println(sl2, len(sl2), cap(sl2))

	sl3 := []int{998: 3}
	fmt.Println(len(sl3), cap(sl3))
	sl4 := append(sl3, 4)
	fmt.Println(len(sl4), cap(sl4))
	//sl5 := append(sl4, 5)
	//fmt.Println(len(sl5), cap(sl5))
}

/*
1.切片是一种基于数组的新的数据结构, 是Go对动态数组的实现 ~ 按需自动増缩(append / 再次切片)
2.切片是一个很小的对象, 对底层数组进行了抽象, 包含3个字段(1.指向底层数组的指针 2.长度 3.容量)
*/
