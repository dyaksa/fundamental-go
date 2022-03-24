package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total
}
func main() {
	slice := []int{10, 50, 60, 32, 30}
	total := sumAll(slice...)
	fmt.Println(total)

	total = sumAll(10, 50, 87, 76)
	fmt.Println(total)
}
