package day04

import (
	"errors"
	"fmt"
	"time"
)

type Err struct {
	info string
	time time.Time
}

func (err Err) Error() string {
	return fmt.Sprint(err.info, err.time)
}

func RetErrs() {
	err := retErrs(3, 2)
	fmt.Println(err)
	fmt.Println("--------------")

	err = retErrs(2, 2)
	fmt.Println(err)
	fmt.Println("--------------")

	err = retErrs(1, 2)
	fmt.Println(err)
}

func retErrs(a, b int) (err error) {
	if a > b {
		err = Err{"自定义error", time.Now()}
	} else if a == b {
		err = errors.New("errors包...")
	} else {
		err = fmt.Errorf("fmt包...")
	}
	return
}
