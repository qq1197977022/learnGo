package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	tk := time.NewTicker(time.Second)
	tm := time.NewTimer(time.Second)
	go ticker(tk)
	go timer(tm)

	wg.Wait()
}

func ticker(tk *time.Ticker) {
	defer wg.Done()
	for {
		<-tk.C
		fmt.Println("间歇定时器")
	}
}

func timer(timer *time.Timer) {
	defer wg.Done()

	for {
		<-timer.C
		fmt.Println("定时器")
	}
}
