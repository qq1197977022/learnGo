package 待分类

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	contents, err := ioutil.ReadFile("C:\\Users\\kc\\Desktop\\10.txt")
	if err != nil {
		fmt.Println(contents)
	} else {
		//fmt.Println(string(contents))
		ioutil.WriteFile("C:\\Users\\kc\\Desktop\\a.txt", contents, os.ModeAppend)
	}
}
