package 待分类

import "fmt"

type notifier1 interface {
	notify()
}

type user1 struct {
	name  string
	email string
}

type admin1 struct {
	name  string
	email string
}

func (u user1) notify() {
	fmt.Printf("%s's <%T> E-mail is %s\n", u.name, u, u.email)
}

func (a admin1) notify() {
	fmt.Printf("%s's <%T> E-mail is %s\n", a.name, a, a.email)
}

func sendNotification1(n notifier1) {
	n.notify()
}
func main() {
	u := user1{"杨一帆", "110@qq.com"}
	a := admin1{"秦晴", "250@qq.com"}

	//多态: 相同代码因类型不同而行为不同
	sendNotification1(u)
	sendNotification1(a)
	sendNotification1(&u)
	sendNotification1(&a)
}
