package day02

import (
	"fmt"
	"log"
)

func Map1() {
	mp1 := map[string]interface{}{"name": "汤姆逊", "age": 30, "gender": true}
	for k, v := range mp1 {
		fmt.Printf("k = %-10s v = %v\n", k, v)
	}

	fmt.Println("------------")
	name := mp1["name"] //非校验式查询
	fmt.Println(name)

	name, ok := mp1["name"] //校验式查询
	if ok {
		fmt.Println(name)
	}

	age, ok := mp1["Age"] //校验式查询
	if ok {
		fmt.Println(age)
	} else {
		log.Println("键值Age不存在", age)
	}
	/*
		1.键值不存在时, 返回元素对应类型的默认零值, 而不是报错
	*/
}
