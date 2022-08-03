package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fo, err := os.Open("task.data")
	if err != nil {
		panic(err)
	}

	defer fo.Close()

	rc := csv.NewReader(fo)
	rc.Comma = ';'

	m, err := rc.ReadAll()
	if err != nil {
		panic(err)
	}

	count := len(m[0]) - 1

	for i := 0; i < count; i++ {
		if m[0][i] == "0" {
			fmt.Println(i + 1)
		}
	}
}
