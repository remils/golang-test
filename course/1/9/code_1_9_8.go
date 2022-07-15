package main

import "fmt"

func main() {
	var x string
	fmt.Scan(&x)
	if (x[0] + x[1] + x[2]) == (x[3] + x[4] + x[5]) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
