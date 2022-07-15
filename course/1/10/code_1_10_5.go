package main

import "fmt"

func main() {
	var x int
	max := 0
	countMax := 1
	for fmt.Scan(&x); x != 0; fmt.Scan(&x) {
		if max == x {
			countMax++
		} else if max < x {
			max = x
			countMax = 1
		}
	}
	fmt.Println(countMax)
}
