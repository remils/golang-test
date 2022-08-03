package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.Replace(s, string('\n'), "", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, ",", ".", -1)
	t := strings.Split(s, ";")
	a, _ := strconv.ParseFloat(t[0], 64)
	b, _ := strconv.ParseFloat(t[1], 64)

	fmt.Printf("%.4f", a/b)
}
