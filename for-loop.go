package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("counter : ", counter)
		counter++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("i : ", i)
	}

	value := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(value); i++ {
		fmt.Println("value : ", value[i])
	}

	persons := make(map[string]string)
	persons["name"] = "Dyaksa"
	persons["title"] = "Programmer"
	for key, value := range persons {
		fmt.Println("key : ", key, "value : ", value)
	}
}
