package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := n / 100
	b := n / 10 % 10
	c := n % 10
	fmt.Print(a + b + c)
}
