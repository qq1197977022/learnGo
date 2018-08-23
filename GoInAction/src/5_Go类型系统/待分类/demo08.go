package 待分类

import "fmt"

type person struct {
	name  string
	email string
}

//使用值接收者声明方法
//notify函数使用值接收者实现了一个方法 ~ 把函数notify和person类型绑定,函数notify变为person类型所属方法
func (u person) notify() {
	//fmt.Printf("Sending person Email To %s<%s>\n", u.name, u.email)
	fmt.Println("\t\t", u)
	u.name = "abc"
	fmt.Println("\t\t", u)
}

//使用指针接收者声明方法
//changeEmail使用指针接收者实现了一个方法 ~ 把函数changeEmail和person类型绑定,函数changeEmail变为person类型所属方法
func (u *person) changeEmail(email string) {
	u.email = email
}

func main() {
	/*bill := person{"Bill", "bill@email.com"}
	bill.notify()

	lisa := &person{"lisa", "lisa@email.com"}
	lisa.notify()

	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()*/

	bill := person{"Bill", "bill@email.com"}
	lisa := &person{"lisa", "lisa@email.com"}

	fmt.Println(bill, lisa)

	bill.notify()
	lisa.notify()

	fmt.Println(bill, lisa)
	fmt.Println("--------------------------------------------")

	fmt.Println(bill, lisa)

	bill.changeEmail("123")
	lisa.changeEmail("456")

	fmt.Println(bill, lisa)
	fmt.Println("===========================================")

	/*fmt.Println(bill, lisa)
	bill.name = "张三"
	lisa.name = "李四"
	fmt.Println(bill, lisa)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++")

	v1 := bill
	v2 := lisa
	fmt.Println(v1, v2)
	fmt.Printf("\t\t%v%v\n", bill, lisa)
	v1.name = "范冰冰"
	v2.name = "李晨"
	fmt.Println(v1, v2)
	fmt.Printf("\t\t%v%v\n", bill, lisa)*/
}

/*
1.方法声明方式: 值/指针接收者
	1.前者使用值副本调用方法,后者使用实际值调用方法
	2.值/指针均可以调用值和指针接收者声明的方法
		1.值调用指针接收者声明的方法时会自动转换为指针
		2.指针调用值接收者声明的方法时会自动转换为值
*/
