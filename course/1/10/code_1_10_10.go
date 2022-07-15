package main

import "fmt"

func main() {
	var s string = ""
	var a, b, c, d int
	fmt.Scan(&a, &b)
	for a > 0 {
		d = a % 10
		a = a / 10
		c = b
		for c > 0 {
			if c%10 == d {
				s = fmt.Sprint(d) + " " + s
			}
			c = c / 10
		}
	}

	fmt.Println(s)
}
