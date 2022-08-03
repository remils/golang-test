package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	var sum int = 0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()

		x, _ := strconv.Atoi(str)

		sum = sum + x
	}

	io.WriteString(os.Stdout, strconv.Itoa(sum))
}
