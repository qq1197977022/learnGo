package day04

import (
	"fmt"
	"time"
)

type person struct {
	name   string
	age    int
	gender bool
}

type myErr struct {
	info string
	time time.Time
	user person
}

func (err myErr) Error() string {
	//return fmt.Sprintln(err.info, err.time, err.user)
	return fmt.Sprintf("%#v", err)
	/*
		疑问:
			1.栈溢出
				1.直接输出err
				2.%+v
				3.%#v不溢出
	*/
}

func CustomErr() {
	retNum, err := customErr()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The input num is %d", retNum)
	}
}

func customErr() (retNum int, err error) {
	num := -1
	fmt.Println("请输入一个正数")
	fmt.Scan(&num)
	if num < 0 {
		err = myErr{fmt.Sprintf("你TM是不是傻...num = %d", num), time.Now(), person{"汤姆", 22, true}}
	} else {
		retNum = num
	}
	return
}
