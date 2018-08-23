package day04

import "fmt"

func RetErr() {
	retNum, err := retErr()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(retNum)
	}
}

func retErr() (retNum int, err error) {
	num := -1
	fmt.Println("请输入一个正数")
	if fmt.Scan(&num); num < 0 {
		err = fmt.Errorf("你TM是不是傻...num = %d", num)
	} else {
		retNum = num
	}
	return
}
