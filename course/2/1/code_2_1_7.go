func minimumFromFour() int {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	array := make([]int, 0, 3)
	array = append(array, b, c, d)
	min := a
	for i := 0; i < 3; i++ {
		if min > array[i] {
			min = array[i]
		}
	}
	return min
}
