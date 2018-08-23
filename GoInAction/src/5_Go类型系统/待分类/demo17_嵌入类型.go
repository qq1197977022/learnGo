package 待分类

import "fmt"

type user4 struct {
	name  string
	email string
}

type admin4 struct {
	user4
	level string
}

type notifier4 interface {
	notify()
}

func (u *user4) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func (adm *admin4) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", adm.name, adm.email)
}

func main() {
	adm := admin4{
		user4: user4{
			name:  "杨一帆",
			email: "1197@qq.com",
		},
		level: "super",
	}

	adm.notify()
	adm.user4.notify()
	sendNotification4(&adm)
	sendNotification4(&adm.user4)
	//外部类型和内部类型均直接实现统一接口,内部类型的实现没有提升(被覆盖)
}

func sendNotification4(n notifier4) {
	n.notify()
}
