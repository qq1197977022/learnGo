package day02

import (
	"flag"
	"fmt"
	"os"
)

/*
疑问:
	1.为什么demo03中函数第二个参数和demo02中同名会报错:  flag redefined
Note:
	1.os.Args持有命令行参数
	2.flag包实现命令行标记解析
		1.命令行标记格式: -flag
*/
func CommandLineArgs() {
	demo01()
	//demo02()
	demo03()
}

func demo01() {
	//测试数据: name 杨一帆 age 21 gender male
	fmt.Println("------方式一: os.Args------")

	fmt.Println(os.Args)
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}

func demo02() {
	//测试数据: -name 杨一帆 -age 21 -gender male
	fmt.Println("------方式二: flag包------")

	namePtr := flag.String("name", "无名氏", "姓名")
	agePtr := flag.Int("age", 0, "年龄")
	genderPtr := flag.String("gender", "male", "性别")

	flag.Parse() //从os.Args[1:]解析命令行参数
	fmt.Printf("%p, %T, %s\n", namePtr, *namePtr, *namePtr)
	fmt.Printf("%p, %T, %d\n", agePtr, *agePtr, *agePtr)
	fmt.Printf("%p, %T, %s\n", genderPtr, *genderPtr, *genderPtr)
}
func demo03() {
	//测试数据: -name 杨一帆 -age 21 -gender male
	fmt.Println("------方式三: flag包------")

	var name string
	var age int
	var gender string

	flag.StringVar(&name, "name", "无名氏", "姓名")
	flag.IntVar(&age, "age", 0, "年龄")
	flag.StringVar(&gender, "gender", "male", "性别")

	flag.Parse()
	fmt.Printf("%p, %T, %s\n", &name, name, name)
	fmt.Printf("%p, %T, %d\n", &age, age, age)
	fmt.Printf("%p, %T, %s\n", &gender, gender, gender)
}
