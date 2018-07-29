package src

import (
	"demo01"
	"fmt"
)

func main() {
	/*
		constant()
		bit()
		complexNum()
		str()
	*/
	//array()
	//HomeWork()

	/*fmt.Printf("%b\n", 22)
	fmt.Println(22<<57)
	fmt.Println(22<<59)*/
	var a int8 = 3
	fmt.Printf("%b\n", a)
	fmt.Println(a << 100)
}
func array() {
	demo01.ArrayLiteral()
}
func str() { //字符串
	demo01.Str()
	demo01.Traverse()
	demo01.Concatenation()
}
func complexNum() { //复数
	demo01.Declar()
	demo01.GetRealORImag()
}
func bit() { //位操作
	fmt.Println("------位运算------")
	demo01.BitSift()   //移位
	demo01.BitAND()    //按位 与
	demo01.BitOR()     //按位 或
	demo01.BitXOR()    //按位 异或
	demo01.BitNOT()    //按位 非
	demo01.BitANDNOT() //按位 Bit Clear
}
func constant() { //常量
	fmt.Println("------常量------")
	demo01.Demo01()
	demo01.Demo02()
}
func HomeWork() {
	demo01.Work01()
}
