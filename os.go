package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Arguments : ")
	fmt.Println(args)

	name, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(name)
}
