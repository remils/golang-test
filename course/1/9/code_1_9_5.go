package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	if x > 0 {
		fmt.Println("Число положительное")
	} else if x < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Ноль")
	}
}
