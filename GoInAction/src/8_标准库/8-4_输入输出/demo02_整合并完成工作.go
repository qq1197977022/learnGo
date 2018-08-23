package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer //声明可变大小的bytes型变量

	b.Write([]byte("hello ")) //追加数据, 根据需要扩容

	fmt.Fprintf(&b, "%q Girl", "Beautiful") //格式化并写入&b

	b.WriteTo(os.Stdout) //把数据写到stdout设备
}

/*
1.该例子展示标准库不同包如何通过支持实现了io.Writer接口类型的值来一起完成工作
2.可以看到接口的优雅及带给语言强大的能力
*/
