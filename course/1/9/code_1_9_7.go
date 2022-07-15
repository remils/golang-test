package main

import "fmt"

func main() {
	var x int
	var a int
	fmt.Scan(&x)
	a = x / 10000
	if a != 0 {
		fmt.Println(a)
	} else {
		a = x % 10000 / 1000
		if a != 0 {
			fmt.Println(a)
		} else {
			a = x % 1000 / 100
			if a != 0 {
				fmt.Println(a)
			} else {
				a = x % 100 / 10
				if a != 0 {
					fmt.Println(a)
				} else {
					fmt.Println(x % 10)
				}
			}
		}
	}
}
