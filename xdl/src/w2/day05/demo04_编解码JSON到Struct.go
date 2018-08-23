package day05

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name    string
	Age     int
	Gender  bool //true: male, false: female
	Hobbies []string
}

func MarshalUnmarshalStruct() {
	marshalStruct()
	unmarshalStruct()
}
func unmarshalStruct() { //解析JSON编码数据
	fmt.Println("------解码JSON------")

	jsonStr := `{"Name":"汤姆","Age":21,"Gender":true,"Hobbies":["足球","篮球","羽毛球"]}`
	byts := bytes.NewBufferString(jsonStr).Bytes()
	var tom person

	if err := json.Unmarshal(byts, &tom); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%v\t %+v\t %T\n", tom, tom, tom)
	}
}

func marshalStruct() { //编码JSON数据
	fmt.Println("------编码JSON------")

	tom := person{"汤姆", 21, true, []string{"足球", "篮球", "羽毛球"}}
	jerry := person{"杰瑞", 18, false, []string{"Windows", "Linux", "Mac"}}
	alex := person{Name: "Alex", Age: 22, Gender: true}
	people := [3]person{tom, jerry, alex}

	if byteSli, err := json.Marshal(tom); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s\n", byteSli)
	}

	if byteSli, err := json.MarshalIndent(tom, "", "\t"); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s\n", byteSli)
	}

	if byteSli, err := json.MarshalIndent(people, "", "\t"); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s\n", byteSli)
	}
}
