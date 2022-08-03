func task2(c chan string, s string) {
	c <- s + " "
	c <- s + " "
	c <- s + " "
	c <- s + " "
	c <- s + " "
}