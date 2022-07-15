package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if c*c == a*a+b*b {
		fmt.Print("Прямоугольный")
	} else {
		fmt.Print("Непрямоугольный")
	}
}
