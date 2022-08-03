package main

import (
	"fmt"
	"sync"
)

func work() {
	fmt.Print("1")
}

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			work()
			defer wg.Done()
		}(wg)
	}

	wg.Wait()
}
