func adding(a string, b string) int64 {
	a1 := []rune(a)
	b1 := []rune(b)
	var a2 string = ""
	for i := range a1 {
		if unicode.IsDigit(a1[i]) {
			a2 = a2 + string(a1[i])
		}
	}
	var b2 string = ""
	for i := range b1 {
		if unicode.IsDigit(b1[i]) {
			b2 = b2 + string(b1[i])
		}
	}
	x, _ := strconv.Atoi(a2)
	y, _ := strconv.Atoi(b2)

	return int64(x) + int64(y)
}
