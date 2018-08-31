package day03

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func CloseChan() {
	wg.Add(2)

	intChan := make(chan int, 3)

	go sender(intChan)
	go receiver(intChan)

	wg.Wait()
}
func sender(intChan chan int) {
	defer wg.Done()

	for i := 5; i > 0; i-- {
		intChan <- i
	}
	close(intChan)
}
func receiver(intChan chan int) {
	defer wg.Done()

	for {
		if num, ok := <-intChan; ok {
			fmt.Println(num, ok)
		} else {
			fmt.Println(num, ok)

			for i := 0; i < 5; i++ {
				num, ok = <-intChan
				fmt.Println(num, ok)
			}

			break
		}
	}
}
