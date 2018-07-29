package main

import (
	"html/template"
	"log"
	"os"
)

type Person struct {
	Name string
	Age  int8
}

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	templateSyntax()
}

func templateSyntax() {
	//t := template.New("NameTemplate") //分配指定名称的HTML模板
	t, err := template.ParseFiles("D:\\workspace\\go\\learnGo\\itcast\\test\\test.html")
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{Name: "杨一帆", Age: 18}
	t.Execute(os.Stdout, p)
}
