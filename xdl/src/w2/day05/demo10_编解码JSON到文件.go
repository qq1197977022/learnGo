package day05

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func MarshalUnmarshalFile() {
	var arr [3]string
	unmarshalFile(arr)
	marshalfile([3]string{"足球", "篮球", "乒乓球"})
}

func unmarshalFile(arr [3]string) { //解码JSON数据
	fmt.Println("------解码JSON数据------")

	fmt.Print("请输入JSON字编码数据: ")
	file := os.Stdin
	defer file.Close()
	dec := json.NewDecoder(file) //创建以标准输入为输入源的Decoder
	//Stdin, Stdout, and Stderr are open Files pointing to the standard input, standard output, and standard error file descriptors.

	if err := dec.Decode(&arr); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(arr)
	}
}

func marshalfile(arr [3]string) { //编码JSON数据
	fmt.Println("------编码JSON数据------")

	file := os.Stdout
	defer file.Close()
	enc := json.NewEncoder(file) //创建以标准输出为输出位置的Encoder
	//Stdin, Stdout, and Stderr are open Files pointing to the standard input, standard output, and standard error file descriptors.
	if err := enc.Encode(&arr); err != nil {
		log.Fatalln(err)
	} /*else {
		fmt.Println(arr)
	}*/
}
