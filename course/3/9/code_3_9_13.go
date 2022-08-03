func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(chan int)

	go func() {
		select {
		case x := <-firstChan:
			output <- x * x
		case x := <-secondChan:
			output <- x * 3
		case <-stopChan:
		}

		defer close(output)
	}()

	return output
}
