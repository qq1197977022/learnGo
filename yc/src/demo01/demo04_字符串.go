package demo01

import "fmt"

func Str() {
	str1 := "abc123区块链"
	fmt.Printf("str1 = %s, len = %d\n", str1, len(str1))
	fmt.Printf("v: str1[0] = %d \t rune: str1[0] = %c\n", str1[0], str1[0])
	/*
		1.String类型表示String值集合, String值是字节序列, 可空
		2.String是不可变类型, 一旦创建, 其内容不可改变
		3.预定义String类型是string ~ defined type
			1. Go允许创建自定义新类型
		4.String类型长度可通过内置函数len()获取, 如果该String是一个常量, 其长度是一个运行时常量
		5.字符串的字节可通过索引访问 eg. str[0]
		6.访问字符串中字节的地址非法
	*/
}
func Traverse() {
	str1 := "abc123区块链"

	fmt.Println("------以byte单位遍历字符串------")
	for i := 0; i < len(str1); i++ {
		fmt.Printf("\t i = %d, v = %d, rune = %c \t %T\n", i, str1[i], str1[i], str1[i])
	}

	fmt.Println("------以rune单位遍历字符串------")
	for i, v := range str1 {
		fmt.Printf("\t i = %d, v = %d, rune = %c \t %T\n", i, v, v, v)
		/*
						十进制			二进制
			str[0] = 	97				‭01100001‬	只读当前字节表示一个Unicode字符
			str1[6] = 	229				11100101	连读3字节表示一个Unicode字符
																					上述体现UTF-8变长编码机制
		*/
	}
}
func Concatenation() {
	str1 := "abc"
	str2 := "123"

	str3 := str1 + str2
	str1 += str2
	fmt.Println("------字符串级联------")
	fmt.Printf("str1 = %s\t str3 = %s\n", str1, str3)
}
