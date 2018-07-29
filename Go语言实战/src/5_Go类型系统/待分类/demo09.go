package 待分类

import "fmt"

func main() {
	var s string
	s = "abc"
	/*
		1.原始类型
			1.分类
				1.数值型
				2.字符串(从技术实现细节上来说也是引用型)
				3.布尔型
			2.Note
				1.值传递
		2.引用类型
			1.分类
				1.切片
				2.映射
				3.通道
				4.接口
				5.函数
			2.Note
				1.引用传递
					1.本质也是值传递 ~ 标头值: 引用类型变量称为标头值(包含用于管理底层数据结构的特定字段)
												eg.切片
													1.指向底层数据结(数组)的指针
													2.长度
													3.容量
					2.传递的是引用类型值的副本,共享的是该引用类型字段中指针所指向的底层数据结构
			3.引用类型本质是基于原始类型的一种新的数据结构
		3.传递类型
			1.值传递 ~ 传递值的副本
			2.址传递 ~ 传递地址
			3.引用传递 ~ 传递引用类型值的副本(标头值 ~ 包含指向底层数据结构的指针)
			4.Note
				1.这三种传递类型本质都是值传递(A.原始类型值, B.地址值, C.引用类型值)
					-----------------------------------------------------
					|~~~传递类型~~~				~~~传递对象~~~			|
					|   值传递					   原始类型值副本			|
					|   址传递					   地址值副本				|
					|   引用传递			           引用类型值副本(标头值)	|
					-----------------------------------------------------
	*/
	fmt.Println(len(s))
}
