package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool

	var noKtpDyaksa NoKtp = "123456789"
	var statusMarried Married = true

	fmt.Println(noKtpDyaksa, statusMarried, "aamiin")
}
