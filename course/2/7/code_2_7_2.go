package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Print(math.Pow((math.Pow(a, 2) + math.Pow(b, 2)), 0.5))
}
