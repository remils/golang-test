package main

import "fmt"

func main() {
	var x int
	var a int
	var b int
	var c int
	fmt.Scan(&x)
	a = x / 100
	b = x % 100 / 10
	c = x % 10
	if a != b && b != c && c != a {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
