for key, _ := range cityPopulation {
	flag := true
	for _, city := range groupCity[100] {
		if city == key {
			flag = false
		}
	}
	if flag == true {
		delete(cityPopulation, key)
	}
}
