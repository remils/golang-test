package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	sa := make([]string, len(r))
	for i := range r {
		sa[i] = string(r[i])
	}
	fmt.Print(strings.Join(sa, "*"))
}
