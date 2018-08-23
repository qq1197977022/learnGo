package pkg2

type user struct {
	Name  string
	Email string
}

type Admin struct {
	user   //内嵌字段 ~ 非公开
	Rights int
}
