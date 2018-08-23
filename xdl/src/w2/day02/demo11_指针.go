package day02

import "fmt"

func Pointer() {
	tom := human{}
	p1(tom)
	p2(&tom)
	fmt.Println(tom)

	num := 3
	p1 := &num
	p2 := &p1
	p3 := &p2
	p4 := &p3
	p5 := &p4
	fmt.Println("------多级指针------")
	fmt.Printf("type: %T\t value: %v\t *var: %v\n", p1, p1, *p1)
	fmt.Printf("type: %T\t value: %v\t *var: %v\n", p2, p2, *p2)
	fmt.Printf("type: %T\t value: %v\t *var: %v\n", p3, p3, *p3)
	fmt.Printf("type: %T\t value: %v\t *var: %v\n", p4, p4, *p4)
	fmt.Printf("type: %T\t value: %v\t *var: %v\n", p5, p5, *p5)
}
func p1(man human) {
	man.age = 33
}
func p2(man *human) {
	man.name = "张三丰"
}
