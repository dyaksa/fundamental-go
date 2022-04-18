package main

import (
	"fmt"

	gosayhello "github.com/dyaksa/go-say-hello"
)

func main() {
	say := gosayhello.SayHello()
	fmt.Println(say)
	name := gosayhello.GetName("Dyaksa")
	fmt.Println(name)
}
