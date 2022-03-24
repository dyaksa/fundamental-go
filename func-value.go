package main

import "fmt"

func getGoodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	goodBye := getGoodbye
	fmt.Println(goodBye("John"))
}
