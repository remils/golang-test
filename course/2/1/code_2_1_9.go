func fibonacci(n int) int {
	f1, f2 := 0, 1
	i := 1
	for {
		if i == n {
			return f2
		}
		f1, f2 = f2, f1+f2
		i++
	}
}

