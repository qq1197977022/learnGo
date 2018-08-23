package 待分类

import "fmt"

type notifier interface {
	notify()
}
type human struct {
	name  string
	email string
}

func (h *human) notify() {
	fmt.Printf("Sending human email to %s<%s>\n", h.name, h.email)
}

func main() {
	h := human{"Tom", "110@qq.com"}
	//sendNotification(h)
	/*
	   # command-line-arguments
	   src\demo12.go:19:18: cannot use h (type human) as type notifier in argument to sendNotification:
	   	human does not implement notifier (notify method has pointer receiver)

	   	h是值类型,其方法集由值接收者声明的方法组成,notify()方法接收者为指针,所有h的方法集不包括notify(),即h没有实现notifier接口
	*/

	sendNotification(&h) //&h是指针类型,其方法集由值接收者和指针接收者声明的方法组成,即&h实现了notifier接口

	/*定义接口接收规则
	   3.方法集
		   1.GO规范方法集解释: T类型的方法集由值接收者声明的方法组成,*T类型的方法集由值/指针接收者声明的方法组成
			   Values						Methods Receivers
				 T								(t T)
				 *T							(t T) and (t *T)
		   2.从接收者类型角度看方法集
			   使用指针(*x)接收者实现接口, 只有*T实现接口
			   使用值(x)接收者实现接口, T和*T都实现接口
			   Methods Receivers			Values
				   (t T)					  T and *T
				   (t *T)						*T
			3.为什么会有这样的规则
				1.编译器有 [自动引用/ 解引用功能 | 自动转换值和指针类型]
				1.但编译器并不是总能自动获取一个值的地址 ~ 参见demo13
				2.所以值的方法集的只包括使用值接收者实现的方法
	*/
}

func sendNotification(h notifier) {
	h.notify()
}
