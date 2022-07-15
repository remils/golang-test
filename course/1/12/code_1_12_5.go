package main

import "fmt"

func main() {
	var workArray [10]uint8

	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}

	var stack [6]uint8

	for i := 0; i < 6; i++ {
		fmt.Scan(&stack[i])

		if i%2 > 0 {
			s := workArray[stack[i]]
			workArray[stack[i]] = workArray[stack[i-1]]
			workArray[stack[i-1]] = s
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Print(workArray[i], " ")
	}

	fmt.Print("type ok")
}
