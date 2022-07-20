package main

import (
	"fmt"
)

func main() {
	var max rune = '0'
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	for i := range r {
		if r[i] > max {
			max = r[i]
		}
	}
	fmt.Print(string(max))
}
