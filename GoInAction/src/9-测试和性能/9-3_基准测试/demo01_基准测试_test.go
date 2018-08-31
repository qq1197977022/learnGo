package main

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	num := 10

	b.ResetTimer()
	/*
		1.把经过的基准时间和内存分配计数器归零, 计时器是否正在运行, 不影响 ~ 一言蔽之: 重置计时器和计数器
		2.跳过初始化代码的执行时间, 以保证测试结果尽量精确
	*/

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", num) //被测函数
	}
}

func BenchmarkFormat(b *testing.B) {
	num := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = strconv.FormatInt(num, 10) //被测函数
	}
}

func BenchmarkItoa(b *testing.B) {
	num := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(num) //被测函数
	}
}

func BenchmarkString(b *testing.B) {
	num := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = string(num) //被测函数
	}
}

/*
1.测试不同函数把字符串转换为整数的效率
	go test
		1.-v: 详情
		2.-bench regexp: 基准测试
		3.-benchtime time.Duration: 测试时间
		4.-benchmem: 内存分配统计
			1.次数: 每次操作在堆上分配内存次数
			2.字节数: 每次操作分配内存字节大小

2.测试结果
	D:\workspace\go\src\9-测试和性能\9-3_基准测试>go test -v -bench=. -benchtime=3s -benchmem
	goos: windows
	goarch: amd64
	BenchmarkSprintf-8      30000000               121 ns/op              16 B/op          2 allocs/op
	BenchmarkFormat-8       1000000000               4.21 ns/op            0 B/op          0 allocs/op
	BenchmarkItoa-8         1000000000               6.74 ns/op            0 B/op          0 allocs/op
	BenchmarkString-8       1000000000               6.54 ns/op            0 B/op          0 allocs/op
	PASS
	ok      _/D_/workspace/go/src/9-测试和性能/9-3_基准测试 23.346s
*/

/*
一.基准测试
	1.目的: 测试代码性能 ~ 同一问题不同方案择优
		1.也可以测试局部某段代码的CPU或内存效率 ~ 该段代码可能会严重影响整个应用程序性能
	2.常见应用
		1.测试不同并发模式	~ 保证最大化系统吞吐量
		2.辅助配置工作池数量	~ ...
	3.测试框架识别测试对象约定
		1.测试文件
			1.必须以 _test.go 结尾
		2.测试函数
			1.签名: 必须公开 ~ 导出
				1.格式: TestXxx ~ 1.以Test开头, 自定义部分Xxx必须首字母大写 ~ 大驼峰
			2.形参: 必须是test.B类型的指针
				1.B类型管理基准测试时间和指定运行迭代次数
			3.返回值: 无
	4.Note
		1.基准测试框架在指定时间内反复调用要测试的函数 ~ 默认测试时间1s
			1.每次调用都会对b.N +1
		2.必须将要进行基准测试的代码都放在循环里, 且循环要使用b.N的值, 否则测试结果不可靠
*/
