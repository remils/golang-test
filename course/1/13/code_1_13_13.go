package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var x int = 1
	for {
		fmt.Print(x, " ")
		x = x * 2
		if x > N {
			break
		}
	}
}
