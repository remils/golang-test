package main

import "fmt"

func main() {
	var x int
	var N int
	fmt.Scan(&N)
	array := make([]int, 0, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&x)
		if x > 0 {
			array = append(array, x)
		}
	}
	fmt.Print(len(array))
}
