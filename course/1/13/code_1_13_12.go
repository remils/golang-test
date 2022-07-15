package main

import "fmt"

func main() {
	var N, x int
	fmt.Scan(&N)
	x = N % 10
	if x == 1 && N != 11 {
		fmt.Print(N, " korova")
	}
	if (N > 4 && N < 21) || (N > 24 && (x > 4 && x < 9) || x == 0) {
		fmt.Print(N, " korov")
	}
	if (N > 1 && N < 5) || (N > 21 && x > 1 && x < 5) {
		fmt.Print(N, " korovy")
	}
}
