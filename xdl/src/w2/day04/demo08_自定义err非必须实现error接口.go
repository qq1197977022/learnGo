package day04

import (
	"fmt"
	"time"
)

//总结: 非必须实现, 但最好与标准库保持一致, 提高兼容性和一致性
type newErr struct {
	info string
	time time.Time
}

func (ne newErr) String() string {
	return fmt.Sprintf("%s\t %s", ne.info, ne.time)
}

func ErrNoImpErrorInterf() {
	defer handleErr()
	panic(newErr{"没有实现error接口的err", time.Now()})
}

/*
疑问:

func ErrNoImpErrorInterf2() error {
	defer handleErr()
	panic(newErr{"没有实现error接口的err", time.Now()})	//若panic必执行, 即使函数声明又返回值, 不返回也合法
}
*/

func handleErr() {
	if err := recover(); err != nil {
		fmt.Printf("---%s---", err)
	}
}
