package main

import "fmt"

type Man struct {
	Name string
}

func Married(man *Man) {
	man.Name = "Mr." + man.Name
}

func main() {
	dyaksa := Man{"Dyaksa"}
	Married(&dyaksa)
	fmt.Println(dyaksa)
}
