package main

import (
	"fmt"

	gosayhello "github.com/dyaksa/go-say-hello/v2"
)

func main() {
	say := gosayhello.SayHelloWithName("Dyaksa")
	fmt.Println(say)
	name := gosayhello.GetName("Dyaksa")
	fmt.Println(name)
}
