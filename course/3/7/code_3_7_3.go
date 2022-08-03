package main

import (
	"fmt"
	"time"
)

func main() {
	var data string

	fmt.Scan(&data)

	date, err := time.Parse(time.RFC3339, data)
	if err != nil {
		panic(err)
	}

	fmt.Print(date.Format(time.UnixDate))
}
