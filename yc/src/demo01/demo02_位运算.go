package demo01

import "fmt"

func BitSift() { //左右移位运算
	var v1 int8 = 3
	fmt.Printf("v1 = %d\t v1 << 1 = %d\n", v1, v1<<1) //左移一位
	/*
		十进制			二进制			左移一位			十进制
		3				00000011		00000110		6
	*/

	var v2 int8 = 3
	fmt.Printf("v2 = %d\t v2 >> 1 = %d\n", v2, v2>>2) //右移一位
	/*
		十进制			二进制			右移一位			十进制
		3				00000011		00000001		1
	*/
	fmt.Printf("Binary: v1 = %b \t v2 = %b\n", v1, v2)
}
func BitAND() {
	fmt.Println("------按位与------")
	var v1 int8 = 10
	var v2 int8 = 5
	var v3 int8 = v1 & v2
	fmt.Printf("Decimal: v1 = %d \t v2 = %d \t v3 = %d\n", v1, v2, v3)
	fmt.Printf("Binary: v1 = %b \t v2 = %b \t v3 = %b\n", v1, v2, v3)
	/*
			二进制			十进制
			00001010		5
			00000101		10
		XOR:00000000		0

		按位与规则:
			1.均为1则为1, 否则为0 ~ 一假则假
			2.双目(二元)运算符
			3.操作数以补码形式参与运算
	*/
}
func BitOR() {
	fmt.Println("------按位或------")
	var v1 int8 = 10
	var v2 int8 = 5
	var v3 int8 = v1 | v2
	fmt.Printf("Decimal: v1 = %d \t v2 = %d \t v3 = %d\n", v1, v2, v3)
	fmt.Printf("Binary: v1 = %b \t v2 = %b \t v3 = %b\n", v1, v2, v3)
	/*
			二进制			十进制
			00001010		5
			00000101		10
		XOR:00001111		15

		异或规则:
			1.有1则1 ~ 一真则真
			2.双目(二元)运算符
			3.操作数以补码形式参与运算
	*/
}
func BitXOR() {
	fmt.Println("------按位异或------")
	var v1 int8 = 10
	var v2 int8 = 5
	var v3 int8 = v1 ^ v2
	fmt.Printf("Decimal: v1 = %d \t v2 = %d \t v3 = %d\n", v1, v2, v3)
	fmt.Printf("Binary: v1 = %b \t v2 = %b \t v3 = %b\n", v1, v2, v3)
	/*
			二进制			十进制
			00001010		5
			00000101		10
		XOR:00001111		15

		异或规则:
			1.相同为0, 不同为1
			2.又称半加法运算 ~ 不进位的二进制加法运算
			3.双目(二元)运算符
			4.操作数以补码形式参与运算
	*/
}
func BitNOT() {
	fmt.Println("------按位非------")
	var v1 int8 = 10
	var v2 int8 = ^v1
	fmt.Printf("Decimal: v1 = %d \t v2 = %d\n", v1, v2)
	fmt.Printf("Binary: v1 = %b \t v2 = %b\n", v1, v2)
	/*
			二进制			十进制
			00001010		10
		NOT:11110101 负数 ---> 取反: 10001010 ---> +1: 10001011 ---> 十进制: -11

		备注:
			1.正数原反补码一致
			2.负数
				1.反码 = 原码取反
				2.补码 = 反码 + 1
			3.参见: https://www.cnblogs.com/everest33Tong/p/6586289.html
	*/
}
func BitANDNOT() {
	fmt.Println("------按位Bit Clear------")
	var v1 int8 = 10
	var v2 int8 = 5
	var v3 int8 = v1 &^ v2
	fmt.Printf("Decimal: v1 = %d \t v2 = %d \t v3 = %d\n", v1, v2, v3)
	fmt.Printf("Binary: v1 = %b \t v2 = %b \t v3 = %b\n", v1, v2, v3)
	/*
		ANDNOT = AND + NOT
	*/
}

/*
1.古老CPU中, 位运算 > 加减运算, 现代CPU中位运算 = 加减运算 > 乘法运算
2.位操作是程序中对位模式按位或二进制数的一元或二元操作
*/
