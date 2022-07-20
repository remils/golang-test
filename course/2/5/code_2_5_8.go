package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text := []rune(input)

	i := 0
	l := len(text) - 1

	for i < l {
		if text[i] != text[l] {
			fmt.Print("Нет")
			return
		}
		i++
		l--
	}

	fmt.Print("Палиндром")
}
