package main

import (
	"encoding/json"
	"fmt"
)

type jsn1 struct {
	/*
		问题: 首字母必须大写, 才能被正确解码, 否则忽略
		解决: 结构体Tag
			1.改名
			2.不显示`jsn1:"-"`
			3.以字符串格式显示`jsn1:",string"`
			4.非导出字段不能添加Tag
		具体: 参见JSON包
	*/
	Company  string   `jsn1:"company"`
	Subjects []string `jsn1:"科目"`
	IsOk     bool     `jsn1:",string"`
	Price    uint16   `jsn1:"-"`
}

func main() {
	jsnStr1 := `{"Company":"itcast","Subjects":["Go","C++","Python","前端"],"IsOk":true,"Price":666}`
	fmt.Println(jsnStr1)
	byts, err := json.MarshalIndent(jsnStr1, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", byts)

	fmt.Println("--------------------------")
	var jsnStr2 jsn1 = jsn1{
		"itcast",
		[]string{"Go", "C++", "Python", "前端"},
		true,
		666,
	}
	fmt.Printf("%+v\n", jsnStr2)
	byts, err = json.MarshalIndent(jsnStr2, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", byts)
}
