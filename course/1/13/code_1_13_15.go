package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var s string = ""
	for N > 0 {
		s = fmt.Sprint(N%2) + s
		N = N / 2
	}
	fmt.Print(s)
}
