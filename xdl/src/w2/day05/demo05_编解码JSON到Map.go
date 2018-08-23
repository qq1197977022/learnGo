package day05

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func MarshalUnmarshalMap() {
	marshalMap()
	unmarshalMap()
}

func marshalMap() { //编码JSON
	fmt.Println("------编码JSON------")

	m := map[string]interface{}{"name": "汤姆", "age": 21, "gender": true, "hobbies": [3]string{"篮球", "足球", "羽毛球"}}
	if bytSli, err := json.Marshal(m); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s\n", bytSli)
	}

	if bytSli, err := json.MarshalIndent(m, "", "\t"); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s\n", bytSli)
	}

}
func unmarshalMap() { //解码JSON
	fmt.Println("------解码JSON------")

	var m map[string]interface{}
	jsonStr := `{"age":21,"gender":true,"hobbies":["篮球","足球","羽毛球"],"name":"汤姆"}`
	bytSli := bytes.NewBufferString(jsonStr).Bytes()
	if err := json.Unmarshal(bytSli, &m); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%v\t %T\n", m, m)
	}
}
