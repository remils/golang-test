package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var a int
	var s int = 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&a)
		if a/100 == 0 && a/10 > 0 && a%8 == 0 {
			s = s + a
		}
	}
	fmt.Println(s)
}
