package main

import (
	"fmt"
	"strconv"
)

func main() {
	pBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pBool)

	fNumber, err := strconv.ParseInt("1000", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fNumber)
}
