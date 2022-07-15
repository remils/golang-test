package main

import "fmt"

func main() {
	var x, sum int
	fmt.Scan(&x)
	sum = 0
	for {
		sum = sum + (x % 10)
		x = x / 10
		if x < 10 {
			sum = sum + x
			if sum < 10 {
				break
			} else {
				x = sum
				sum = 0
			}
		}
	}

	fmt.Print(sum)
}
