package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("Dyaksa Jauharuddin", "Dyaksa"))
	fmt.Println(strings.Contains("Dyaksa", "Aksa"))

	fmt.Println(strings.Split("Dyaksa Jauharuddin", " "))

	fmt.Println(strings.ToUpper("dyaksa jauharuddin"))
	fmt.Println(strings.ToLower("DYAKSA JAUHARUDDIN"))
	fmt.Println(strings.ToTitle("dyaksa jauharuddin"))

	fmt.Println(strings.Trim("  Dyaksa Jauharuddin  ", " "))

	fmt.Println(strings.ReplaceAll("Dyaksa Jauharuddin", "Dyaksa", "Aksa"))
}
