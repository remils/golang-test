package main

import "fmt"

func main() {
	var N, x, min, count int
	fmt.Scan(&N, &min)
	count = 1
	for i := 0; i < N; i++ {
		fmt.Scan(&x)
		if x == min {
			count++
		} else if x < min {
			min = x
			count = 1
		}
	}
	fmt.Print(count)
}
