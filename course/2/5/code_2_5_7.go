package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text := []rune(input)

	if unicode.IsUpper(text[0]) && text[len(text)-1] == '.' {
		fmt.Print("Right")
	} else {
		fmt.Print("Wrong")
	}
}
