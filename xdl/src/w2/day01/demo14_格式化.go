package day01

import "fmt"

func Format() {
	/*
		宽度 和 对齐方式: 默认最小宽度 + 左对齐
	*/
	fmt.Println("------宽度 和 对齐方式------")
	fmt.Printf("-%s-\n", "中华人民共和国")
	fmt.Printf("-%20s-\n", "中华人民共和国")
	fmt.Printf("-%-20s-\n", "中华人民共和国")
	/*
		格式化顺序
		1.默认从左到右一一对应格式化
	*/
	fmt.Println("------格式化顺序------")
	fmt.Printf("%[3]s\t %[1]s \t%s\n", "中华", "人民", "共和国")
	fmt.Printf("%[1]s\t %[1]s \t%s\n", "中华", "人民", "共和国")
	fmt.Printf("%[3]s\t %[3]s \t%[3]s\n", "中华", "人民", "共和国")
}
