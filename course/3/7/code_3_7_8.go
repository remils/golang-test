package main

import (
	"fmt"
	"time"
)

// вам это понадобится
const now = 1589570165

func main() {
	var min, sec string

	fmt.Scanf("%s мин. %s сек.", &min, &sec)

	d, _ := time.ParseDuration(fmt.Sprintf("%sm%ss", min, sec))

	fmt.Println(time.Unix(now, 0).Add(d).UTC().Format(time.UnixDate))
}
