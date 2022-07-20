func sumInt(array ...int) (int, int) {
	l := len(array)
	s := 0
	for i := 0; i < l; i++ {
		s = s + array[i]
	}
	return l, s
}
