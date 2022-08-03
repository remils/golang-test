// создать функцию проверки
func checkValue(value interface{}) bool {
	switch v := value.(type) {
	case float64:
		return true
	default:
		fmt.Printf("value=%v: %T", v, v)
		return false
	}
}

func main() {
	// вставить в main
	if checkValue(value1) && checkValue(value2) {
		switch operation.(string) {
		case "+":
			fmt.Printf("%.4f", value1.(float64)+value2.(float64))
		case "-":
			fmt.Printf("%.4f", value1.(float64)-value2.(float64))
		case "*":
			fmt.Printf("%.4f", value1.(float64)*value2.(float64))
		case "/":
			fmt.Printf("%.4f", value1.(float64)/value2.(float64))
		default:
			fmt.Print("неизвестная операция")
		}
	}
}
