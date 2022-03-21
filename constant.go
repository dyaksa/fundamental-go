package main

import "fmt"

func main() {
	const firstName = "Dyaksa"
	const lastName string = "Jauharuddin"
	const age = 27

	fmt.Println("My name is", firstName, lastName, "Age", age)

	const (
		friendName string = "Alfa"
		friendAge         = 25
	)

	fmt.Println("My friend's name is", friendName, "Age", friendAge)
}
