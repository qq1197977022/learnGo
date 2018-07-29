package 待分类

import "fmt"

func main() {
	var sl1 = make([][]int, 2, 10)
	fmt.Printf("%v\t%v\t%v\t%p\n", sl1, len(sl1), cap(sl1), &sl1)

	sl1 = [][]int{{1}, {10, 20}}
	fmt.Printf("%v\t%v\t%v\t%p\n", sl1, len(sl1), cap(sl1), &sl1)

	fmt.Println("------------")
	sl1 = append(sl1, []int{100})
	fmt.Printf("%v\t%v\t%v\t%p\n", sl1, len(sl1), cap(sl1), &sl1)

	/*sl2 := append(sl1, []int{100})
	fmt.Printf("%v\t%v\t%v\t%p\n", sl2, len(sl2), cap(sl2), &sl2)*/
	/*为什么地址不同?
	因为&sl1, &sl2获取的切片变量的地址值,而不是切片这种数据结构指针所指向的底层数组的地址值
	*/

	sl3 := append(sl1[0], 2)
	sl3 = append(sl3, 3)
	sl3 = append(sl3, 4)

	fmt.Println(len(sl3), cap(sl1))
	fmt.Println(sl1, sl3, len(sl1[0]))
}
