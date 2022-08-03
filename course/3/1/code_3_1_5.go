var cache = map[int]int{}

var x int

for i := 0; i < 10; i++ {
	fmt.Scan(&x)
	_, inMap := cache[x]
	if !inMap {
		cache[x] = work(x)
	}
	fmt.Print(cache[x], " ")
}
