package 待分类

import "fmt"

func main() {
	fmt.Println("Go提供3种数据结构-数组-切片-映射\t", "数组是切片和映射的基础数据结构", "\t切片和映射的底层是数组")
	//声明数组
	var arr1 = [5]int{}

	//数组字面
	var arr2 = [5]int{1, 2, 3, 4, 5}
	var arr3 = [...]int{1, 2, 3, 4, 5} //...等于最大元素索引+1
	var arr4 = [5]int{0: 1, 3: 4}      //关键字元素
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println("------------------")

	//声明元素为int型指针的数组
	var arr5 = [5]*int{}
	fmt.Println(arr5)
	arr5[0] = new(int) //new(T)分配内存并初始化T,然后返回分配内存的地址
	arr5[3] = new(int)
	*arr5[3] = 4 //*操作符访问指针指向的内存中的值
	fmt.Println(arr5)

	fmt.Println(*arr5[0], *arr5[3]) //*操作符访问指针指向的内存
	fmt.Println("------------------")

	//数组是值传递
	var arr6 = [3]string{"A"}
	var arr7 = [3]string{"a"}
	fmt.Printf("%#v\t\t%v\n", arr6, arr7)

	arr7 = arr6
	fmt.Printf("%#v\t\t%v\n", arr6, arr7)
	arr6[0] = "B"
	fmt.Printf("%#v\t\t%v\n", arr6, arr7)
	fmt.Println("------------------")

	//多维数组
	var arr8 = [2][3]int{}
	fmt.Println(arr8)
}
