package main

import "fmt"

func sayitHello() (firstName string, lastName string) {
	firstName = "Dyaksa"
	lastName = "Jauharuddin"
	return
}
func main() {
	firstName, lastName := sayitHello()
	fmt.Println(firstName)
	fmt.Println(lastName)
}
