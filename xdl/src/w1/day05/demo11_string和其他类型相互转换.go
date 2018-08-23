package day05

import (
	"fmt"
	"log"
	"strconv"
)

func Parse() { //将字符串解析为指定类型
	str := "true"
	v2, err := strconv.ParseBool(str)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Parse: %T\t%v ---> %T\t%v\n", str, str, v2, v2)
}
func Format() { //将指定类型格式化为字符串
	b := true
	v1 := strconv.FormatBool(b)
	fmt.Printf("format: %T\t%v ---> %T\t%v\n", b, b, v1, v1)
}
