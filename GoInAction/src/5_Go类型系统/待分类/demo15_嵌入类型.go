package 待分类

import (
	"fmt"
)

type user2 struct {
	name  string
	email string
}

func (u *user2) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin2 struct {
	user2 //嵌入字段
	level string
}

func main() {
	adm := admin2{
		user2: user2{
			name:  "杨一帆",
			email: "1197@qq.com",
		},
		level: "super",
	}

	adm.user2.notify() //通过内嵌字段,访问内嵌字段包含的标识符
	adm.notify()       //内嵌字段标识符自动提升至外部类型
}

/*
1.嵌入类型: 在新的结构体类型中直接声明已有类型,称之为新的外部类型的内部类型
2.内部类型的标识符会被提升到外部类型上,成为外部类型的一部分 ~ 即外部类型组合了内部类型的属性,且可添加新的字段和方法;外部类型也可以覆盖内部标识符的字段或方法 ~ 即扩展或修改已有类型
*/
/*
结构体字段:
	1.显示命名
	2.隐式命名 ~ 嵌入字段
*/
