package main

import "fmt"

func main() {
	var N, i int
	fmt.Scan(&N)
	if N == 0 {
		fmt.Print(0)
	} else {
		fib_prev, fib_next := 0, 1
		i = 1
		for {
			if fib_next > N {
				fmt.Print(-1)
				break
			}

			if fib_next == N {
				fmt.Print(i)
				break
			}

			fib_prev, fib_next = fib_next, fib_prev+fib_next
			i++
		}
	}
}
