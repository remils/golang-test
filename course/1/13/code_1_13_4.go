package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	h := t / 3600
	m := t % 3600 / 60
	fmt.Println("It is", h, "hours", m, "minutes.")
}
