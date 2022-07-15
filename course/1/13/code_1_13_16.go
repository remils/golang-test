package main

import "fmt"

func main() {
	var N, x int
	fmt.Scan(&N, &x)
	var s string = ""
	for N > 0 {
		if N%10 != x {
			s = fmt.Sprint(N%10) + s
		}
		N = N / 10
	}
	fmt.Print(s)
}
