package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x, &p, &y)
	var c int = 0
	for ; x < y; c++ {
		x = x * (100 + p) / 100
	}
	fmt.Println(c)
}
