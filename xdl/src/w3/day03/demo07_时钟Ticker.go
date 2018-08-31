package day03

import (
	"fmt"
	"time"
)

func Ticker() {
	ticker()
}

func ticker() {
	tickerPtr := time.NewTicker(time.Second) //创建时钟
	for i := 0; i < 8; i++ {
		timeNow := <-tickerPtr.C
		fmt.Printf("%d\t %v\n", i, timeNow)
	}
}
