package main

import "fmt"

func main() {
	var name string

	name = "Dyaksa"
	fmt.Println(name)

	name = "Dyaksa Jauharuddin"
	fmt.Println(name)

	var friendName = "Alfa"
	fmt.Println(friendName)

	var age = 25
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "Dyaksa"
		lastName  = "Jauharuddin"
	)

	fmt.Println("My name is", firstName, lastName)
}
