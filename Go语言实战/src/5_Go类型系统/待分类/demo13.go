package 待分类

import "fmt"

type duration int //自定义一个基于int类型的类型

func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

func main() {
	duration(42).pretty()
	/*~~demo12例子~~
	# command-line-arguments
	src\tempFile.go:12:14: cannot call pointer method on duration(42)
	src\tempFile.go:12:14: cannot take the address of duration(42)
	*/
}
