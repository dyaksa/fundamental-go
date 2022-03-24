package main

import "fmt"

func sayHelloTo(firstName string, lastName string) string {
	return "hello " + firstName + " " + lastName
}

func main() {
	res := sayHelloTo("Dyaksa", "Jauharuddin")
	fmt.Println(res)
}
