package 待分类

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer

	b.Write([]byte("hello"))

	b.Write([]byte("world"))
	fmt.Printf("%+v\n", b)

	s1 := b.Bytes()

	fmt.Printf("%v\n", string(s1))

	fmt.Println(b.String())

	/*fmt.Fprintf(&b," %v", "world")
	io.Copy(os.Stdout, &b)*/
}

/*
1.接口作用: 用于定义行为 ~ 声明方法
	1.行为(方法)不由接口直接实现,而由用户定义类型(实体类型)实现,此时接口类型变量可以接收(存储)用户类型的值
1.多态: 相同代码因类型不同而行为不同
2.若某类型实现某个接口,则所有使用该接口的地方均支持该类型的值
*/
