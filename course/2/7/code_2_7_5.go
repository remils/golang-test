package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	sr := ""
	for i := range r {
		x, _ := strconv.Atoi(string(r[i]))
		sr = sr + fmt.Sprint(x*x)
	}
	fmt.Print(sr)
}
