package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	byts, err := ioutil.ReadFile("D:\\workspace\\go\\learnGo\\itcast\\src\\01_20H\\day05_异常_文本文件处理\\html\\day21.html")
	if err != nil {
		fmt.Print(err)
	}
	str := bytes.NewBuffer(byts).String()

	reg := regexp.MustCompile("<p>(.*)</p>")
	strSli := reg.FindAllStringSubmatch(str, -1)
	fmt.Println(strSli, len(strSli))

	f, err := os.Create("D:\\workspace\\go\\learnGo\\itcast\\src\\01_20H\\day05_异常_文本文件处理\\html\\day21.txt")
	if err != nil {
		fmt.Println(err)
	}
	for _, v1 := range strSli {
		f.WriteString(fmt.Sprintf("%s\n", v1[1]))
	}

}
