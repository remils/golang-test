package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	c := a*a + b*b
	fmt.Println(c)
}
