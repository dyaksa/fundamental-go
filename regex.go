package main

import (
	"fmt"
	"regexp"
)

func main() {

	regex := regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))

	fmt.Println(regex.FindAllString("eko eka edo endah", 10))

}
