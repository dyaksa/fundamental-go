package main

import "fmt"

func main() {
	var name string = "Dyaksa"

	if name == "Dyaksa" {
		fmt.Println("helo dyaksa")
	} else if name == "winandi" {
		fmt.Println("Hello winandi")
	} else {
		fmt.Println("Boleh kenalan")
	}

	if length := len(name); length > 5 {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("nama terlalu pendek")
	}
}
