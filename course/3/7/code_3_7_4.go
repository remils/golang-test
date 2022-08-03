package main

import (
	"fmt"
	"time"
)

func main() {
	var timeFormat string = "2006-01-02 15:04:05"
	var d, t string
	fmt.Scan(&d, &t)

	data := fmt.Sprintf("%s %s", d, t)

	date, err := time.Parse(timeFormat, data)
	if err != nil {
		panic(err)
	}

	if date.Hour() > 13 {
		date = date.AddDate(0, 0, 1)
	}

	fmt.Print(date.Format(timeFormat))
}
