func removeDuplicates(inputStream, outputStream chan string) {
	temp := ""
	for value := range inputStream {
		if temp != value {
			temp = value
			outputStream <- value
		}
	}

	close(outputStream)
}
