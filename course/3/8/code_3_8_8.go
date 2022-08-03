func task(c chan int, N int) {
	c <- N + 1
}