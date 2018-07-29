package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type jsn2 struct {
	Company  string
	Subjects [4]string
	IsOk     bool
	Price    int
}

func main() {
	jsnStr := `{"Company":"itcast","Subjects":["Go","C++","Python","前端"],"IsOk":true,"Price":666}`
	var jsn = &jsn2{}

	bytSli := bytes.NewBufferString(jsnStr).Bytes() //字符串转字节切片
	err := json.Unmarshal(bytSli, jsn)              //解码JSON数据, 到结构体
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(jsn)
	fmt.Println(*jsn)
	fmt.Printf("%+v\n", jsn)
	fmt.Printf("%+v\n", *jsn)
	fmt.Println("====================================")
	/*
		1.结构体有几个对应字段, 则解析几条数据
		待细分总结...
	*/

	bytSli, err = json.MarshalIndent(jsn, "", "	") //编码结构体数据, 到JSON数据
	if err != nil {
		log.Fatalln(err)
	}
	jsnStr = bytes.NewBuffer(bytSli).String() //字节切片转字符串
	fmt.Println(jsnStr)
}
