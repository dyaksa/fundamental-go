package main

import "fmt"

func main() {
	value := 0

	increment := func() {
		fmt.Println("increment")
		value++
	}

	increment()
	fmt.Println(value)
}
