package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)

	if len(r) < 5 {
		fmt.Print("Wrong password")
		return
	}

	for i := range r {
		if !unicode.Is(unicode.Latin, r[i]) && !unicode.IsDigit(r[i]) {
			fmt.Print("Wrong password")
			return
		}
	}

	fmt.Print("Ok")
}
