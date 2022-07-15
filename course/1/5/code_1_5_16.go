package main

import "fmt"

func main() {
	var deg int
	fmt.Scan(&deg)
	h := deg / 30
	m := deg % 30 * 2
	fmt.Println("It is", h, "hours", m, "minutes.")
}
