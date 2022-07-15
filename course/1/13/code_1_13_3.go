package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)
	fmt.Print(string(n[2]))
	fmt.Print(string(n[1]))
	fmt.Print(string(n[0]))
}
