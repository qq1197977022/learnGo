package day04

import (
	"fmt"
	"reflect"
)

func GetMethodInfo() {
	tom := Human{"汤姆", 22}
	/*
		疑问:
			1.tom是*Human类型
				1.NumField()/ Field()报错非结构体类型?
			2.卧槽...
				1.(*Human)Eat	(Human)Run
					1.TypeOf(&tom)	TypeOf(tom) 上下NumMethod均返回2
					2.TypeOf(tom), TypeOf(&tom) 上下NumMethod均返回1
				2.(*Human)Eat	(*Human)Run
					1.TypOf(tom) NumMethod()返回0
					2.TypOf(&tom) NumMethod()返回2
				3.(Human)Eat	(Human)Run
					1.TypeOf(&tom) NumMethod返回2
					2.TypeOf(tom) NumMethod返回2
	*/
	typ1 := reflect.TypeOf(tom)
	fmt.Println(typ1)
	fmt.Println(typ1.NumMethod())

	fmt.Println("-------------------")
	typ2 := reflect.TypeOf(&tom)
	fmt.Println(typ2)
	fmt.Println(typ1.NumMethod())
	//fmt.Println(typ2.NumField())
	//fmt.Println(typ2.Field(0).Anonymous)
}
func (h Human) Eat(goods string) {
	fmt.Printf("%s 吃 %s\n", h.name, goods)
}
func (h Human) Run(speed int) {
	fmt.Printf("%s: 心情是自由自在, 速度是%d迈...\n", h.name, speed)
}
