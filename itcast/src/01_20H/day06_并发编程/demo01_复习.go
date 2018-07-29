package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	byts, err := ioutil.ReadFile("D:\\workspace\\go\\learnGo\\itcast\\src\\01_20H\\day05_异常_文本文件处理\\html\\day21.html")
	if err != nil {
		fmt.Println(err)
	}
	buf := bytes.NewBuffer(byts)
	str := buf.String()
	fmt.Println(str)

	re := regexp.MustCompile("<\\w>(.*)</\\w>")
	fmt.Println(re.FindAllString(str, -1))
	fmt.Println(re.FindAllStringSubmatch(str, -1))

}
