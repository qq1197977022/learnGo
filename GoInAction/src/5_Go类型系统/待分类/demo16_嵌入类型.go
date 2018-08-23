package 待分类

import (
	"fmt"
)

type user3 struct {
	name  string
	email string
}

type admin3 struct {
	user3
	level string
}

type notifier3 interface {
	notify()
}

func (u *user3) notify() { //通过指针接收者实现接口,对应的只有该类型的指针类型才能实现该接口
	fmt.Printf("Sending user3 email to %s<%s>\n", u.name, u.email)
}

func main() {
	adm1 := &admin3{
		user3: user3{
			name:  "杨一帆",
			email: "1197@qq.com",
		},
		level: "super",
	}

	adm2 := admin3{
		user3: user3{
			name:  "秦晴",
			email: "77520@qq.com",
		},
		level: "low",
	}

	sendNotification3(adm1) //由于内部类型的提升,外部类型间接实现notifier接口
	sendNotification3(&adm2)

	fmt.Println("---------------")

	sendNotification3(&adm1.user3) //adm1.user3值 ---> &adm1.user3址
	sendNotification3(&adm2.user3) //同上

	//fmt.Println(&adm1)
	//fmt.Printf("%p", adm1)
}

func sendNotification3(n notifier3) {
	n.notify()
}
