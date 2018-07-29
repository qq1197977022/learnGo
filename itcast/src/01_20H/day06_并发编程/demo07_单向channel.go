package main

import (
	"fmt"
	"log"
	"sync"
)

var wg1 sync.WaitGroup

func init() {
	log.SetFlags(log.Lshortfile)
}

func producer1(sch chan<- int) {
	defer wg1.Done()

	for i := 0; i < 10; i++ {
		sch <- i
		fmt.Printf("ch<------生产 %d\n", i)
	}
	close(sch)
}
func consumer1(rch <-chan int) {
	defer wg1.Done()
	for v := range rch {
		fmt.Printf("消费 %d\n", v)
	}
}

func main() {
	wg1.Add(2)
	ch := make(chan int)

	go producer1(ch)
	go consumer1(ch)

	wg1.Wait()
}
