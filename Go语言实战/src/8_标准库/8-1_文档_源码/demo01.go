package main

func main() {

}

/*
1.标准库源码对Go工具非常重要, (godoc, gocode, go build ...etc)等Go工具均需要读取标准库源码才能完成工作
2.标准库源码都是经过预编译的, 编译后的文件称为归档文件, 以.a为扩展名
	1.归档文件是特殊的GO静态库文件, 由Go构建工具创建, 并在编译和链接最终程序时被使用
	2.归档文件可以提高构建速度
	3.构建过程中, 无法指定这些文件, Go工具链知道何时使用归档文件 / 从源码重新构建
*/
