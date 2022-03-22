package main

import "fmt"

func main() {
	var name string = "Dyaksa"

	switch name {
	case "Dyaksa":
		fmt.Println("Hello Dyaksa")
	case "Tifani":
		fmt.Println("Helo tifani")
	default:
		fmt.Println("halo boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah sesuai")
	}

	length := len(name)
	switch {
	case length > 5:
		fmt.Println("nama cukup panjang")
	case length > 10:
		fmt.Println("nama sangat panjang")
	default:
		fmt.Println("name sudah sesuai")
	}
}
