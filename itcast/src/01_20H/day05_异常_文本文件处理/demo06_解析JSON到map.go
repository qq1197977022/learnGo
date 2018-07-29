package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func main() {
	jsnStr := `{"Company":"itcast","Subjects":["Go","C++","Python","前端"],"IsOk":true,"Price":666}`
	m := make(map[string]interface{}, 4)

	bytSli := bytes.NewBufferString(jsnStr).Bytes()

	err := json.Unmarshal(bytSli, &m)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(m)

	fmt.Println(reflect.TypeOf(m["Subjects"]).Kind())

}
