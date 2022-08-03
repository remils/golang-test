fn := func(x uint) uint {
	str := strconv.Itoa(int(x))

	newStr := ""

	for i := range str {
		value := string(str[i])
		if value == "2" || value == "4" || value == "6" || value == "8" {
			newStr = newStr + value
		}
	}

	t, _ := strconv.Atoi(newStr)

	if t == 0 {
		return 100
	}

	return uint(t)
}
