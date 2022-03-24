package main

import "fmt"

type Filter func(string) string

func sayHelloFilter(name string, filter Filter) string {
	filtered := "hello " + filter(name)
	return filtered
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	filter := sayHelloFilter("Dyaksa", spamFilter)
	fmt.Println(filter)
	filter = sayHelloFilter("anjing", spamFilter)
	fmt.Println(filter)
}
