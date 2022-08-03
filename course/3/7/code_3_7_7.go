package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	formatDateTime := "02.01.2006 15:04:05"

	scanner := csv.NewReader(os.Stdin)

	t, _ := scanner.ReadAll()

	t1, _ := time.Parse(formatDateTime, t[0][0])
	t2, _ := time.Parse(formatDateTime, t[0][1])

	if t1.After(t2) {
		fmt.Println(t1.Sub(t2).Round(time.Second))
	} else {
		fmt.Println(t2.Sub(t1).Round(time.Second))
	}
}
