package pkg2

type alertCounter int

func New(i int) alertCounter { //工厂函数
	return alertCounter(i)
}
