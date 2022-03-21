package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Dyaksa"
	names[1] = "Tifani"
	names[2] = "Widodo"

	fmt.Println(names)

	var values = [3]int{
		1,
		2,
		3,
	}

	fmt.Println(values)
	fmt.Println(len(values))
}
