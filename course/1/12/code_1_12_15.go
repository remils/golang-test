package main

import "fmt"

func main() {
	var x int
	var N int
	fmt.Scan(&N)
	array := make([]int, 0, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&x)
		if i%2 == 0 {
			array = append(array, x)
		}
	}
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i], " ")
	}
}
