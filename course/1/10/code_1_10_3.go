package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var s int = 0
	for i := a; i <= b; i++ {
		s = s + i
	}
	fmt.Println(s)
}
