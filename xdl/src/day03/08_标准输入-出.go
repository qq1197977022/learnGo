package day03

import "fmt"

func StandIO() {
	var name string
	var age int

	fmt.Println("请输入")
	//fmt.Scanf("name = %s age = %d", &name, &age)	//以指定格式输入
	/*
		fmt.Scanf("%s+%d", &name, &age)
			Bug: 所有输入都会识别为字符串
			Debug: %d+%s
	*/
	//fmt.Scanln(&name, &age)
	fmt.Scan(&name, &age)
	fmt.Printf("name = %s\t age = %d", name, age)
}
