package day05

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

var sli []string

func MarshalUnmashalSli() {
	marshalSli()
	unmarshalSli()
}
func marshalSli() { //编码JSON
	fmt.Println("------编码JSON------")

	sli = []string{"足球", "篮球", "乒乓球"}
	if bytSli, err := json.Marshal(sli); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s\n", bytSli)
	}
}
func unmarshalSli() { //解码JSON
	fmt.Println("------解码JSON------")

	jsonStr := `["足球","篮球","乒乓球"]`
	bytSli := bytes.NewBufferString(jsonStr).Bytes()
	if err := json.Unmarshal(bytSli, &sli); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%v\t %T\n", sli, sli)
	}
}
