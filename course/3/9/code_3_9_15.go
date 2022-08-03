package main

import (
	"fmt"
)

const N = 10

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int) {
		done := make(chan bool, n*2)
		results := make([]*int, n*2)

		go func() {
			nDone := 0
			for <-done {
				for i := nDone; i < n; i++ {
					if results[i] != nil && results[i+n] != nil {
						out <- (*results[i] + *results[i+n])
						if nDone++; nDone == n {
							return
						}
					} else {
						break
					}
				}
			}
		}()
		input := func(ch <-chan int, results []*int) {
			for i := 0; i < n; i++ {
				x := <-ch
				go func(i int, x int) {
					result := f(x)
					results[i] = &result
					done <- true
				}(i, x)
			}
		}
		go input(in1, results[:n])
		go input(in2, results[n:])
	}(fn, in1, in2, out)
}

func main() {
	in1, in2, out := make(chan int, 1), make(chan int, 1), make(chan int, N)

	merge2Channels(some, in1, in2, out, N)

	for i := 0; i < N; i++ {
		in1 <- i
	}

	for i := 0; i < N; i++ {
		in2 <- i
	}

	for i := 0; i < N; i++ {
		fmt.Println(<-out)
	}
}

func some(val int) int {
	return val
}
