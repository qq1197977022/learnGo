package day05

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func MarshalUnmashalArr() {
	var arr [3]string
	marshalArr(arr)
	unmarshalArr(arr)
}
func marshalArr(arr [3]string) { //编码JSON
	fmt.Println("------编码JSON------")

	arr = [3]string{"足球", "篮球", "乒乓球"}
	if bytSli, err := json.Marshal(arr); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%s\n", bytSli)
	}
}
func unmarshalArr(arr [3]string) { //解码JSON
	fmt.Println("------解码JSON------")

	jsonStr := `["足球","篮球","乒乓球"]`
	bytSli := bytes.NewBufferString(jsonStr).Bytes()
	if err := json.Unmarshal(bytSli, &arr); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%v\t %T\n", arr, arr)
	}
}
