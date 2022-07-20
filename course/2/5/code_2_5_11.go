package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, r string
	fmt.Scan(&s)
	r = ""
	l := len(s)
	for i := 0; i < l; i++ {
		if strings.Count(s, string(s[i])) > 1 {
			continue
		}
		r = r + string(s[i])
	}
	fmt.Print(r)
}
