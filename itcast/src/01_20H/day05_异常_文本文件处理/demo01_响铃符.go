package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%c", '\a')
	time.Sleep(3 * time.Second)
	fmt.Printf("%s", "\a")
	time.Sleep(3 * time.Second)
	fmt.Print("\a")
	time.Sleep(3 * time.Second)
	fmt.Print('\a')
	//IDE直接运行, 不响铃, 必须在控制台运行才响铃
}
