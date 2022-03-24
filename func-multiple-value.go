package main

import "fmt"

func getName() (string, string) {
	return "dyaksa", "jauharuddin"
}

func main() {
	firstName, lastName := getName()
	fmt.Println(firstName, lastName)
}
