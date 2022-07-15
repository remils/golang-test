package main

import "fmt"

func main() {
	var x int
	for fmt.Scan(&x); x < 101; fmt.Scan(&x) {
		if x < 10 {
			continue
		}
		fmt.Println(x)
	}
}
