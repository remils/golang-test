package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	count int
}

func (b Battery) String() string {
	str := ""

	for i := 10; i > b.count; i-- {
		str = str + " "
	}

	for i := b.count; i > 0; i-- {
		str = str + "X"
	}

	return fmt.Sprintf("[%v]", str)
}

func main() {
	var x string

	fmt.Scan(&x)

	count := strings.Count(x, "1")

	batteryForTest := Battery{count}

	fmt.Print(batteryForTest)
}
