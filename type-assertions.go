package main

import "fmt"

func random() interface{} {
	return false
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("result :", value)
	case int:
		fmt.Println("result : ", value)
	default:
		fmt.Println("uknown")
	}
}
