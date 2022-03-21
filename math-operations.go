package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a int32 = 20
	var b int32 = 10
	var c int32 = a * b

	fmt.Println(c)

	var ab = 10
	ab += 10
	fmt.Println(ab)

	ab++
	fmt.Println(ab)
}
