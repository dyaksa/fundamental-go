package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return false
	} else {
		return "hello world"
	}
}

func main() {
	say := Ups(1)
	fmt.Println(say)
}
