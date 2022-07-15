package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	s := a + b
	r := s / 2
	fmt.Print(r)
}
