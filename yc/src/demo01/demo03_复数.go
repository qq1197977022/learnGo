package demo01

import "fmt"

func Declar() {
	fmt.Println("------复数声明------")
	c1 := 1.3 + 2.4i //默认Complex128, 即虚部, 实部 float64
	var c2 complex64 = 3.5 + 2.7i
	c3 := complex(3.5, 2.1) //内置函数构造复数
	fmt.Printf("c1 = %g \t %T\n", c1, c1)
	fmt.Printf("c2 = %g \t %T\n", c2, c2)
	fmt.Printf("c3 = %g \t %T\n", c3, c3)
}
func GetRealORImag() {
	fmt.Println("------获取 实部 和 虚部------")
	c1 := 1.3 + 2.4i
	rea := real(c1)
	ima := imag(c1)
	fmt.Printf("c1 = %g \n\t rea = %f\n\t ima = %f", c1, rea, ima)
}
