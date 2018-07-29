package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	/*open: //label
	tempFile, err := os.Open("D:\\workspace\\备忘.txt")
	defer tempFile.Close()

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%T, %p", tempFile, tempFile)
	goto open //无条件跳转*/

	/*byt := make([]byte, 32)
	n, err := tempFile.Read(byt)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%d\t%s", n, byt)*/

	/*byt, err := ioutil.ReadFile("D:\\workspace\\备忘.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s", byt)*/

	//bufReader := bufio.NewReader(tempFile)
	/*byt, err := bufReader.ReadByte()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%d", byt)*/

	/*byts, b, err := bufReader.ReadLine()
	fmt.Printf("%s %v %v\n", byts, b, err)*/

	/*byts, err := bufReader.ReadSlice('M')
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s", byts)*/

	/*str, err := bufReader.ReadString('M')
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(str)*/

	m := f
	m()
	fmt.Printf("m = %p\n", m)
	fmt.Printf("f = %p\n", f)
}
func f() {
	fmt.Println("函数类型", f)
}
