package main

import (
	"flag"
	"fmt"
)

func main() {

	var host *string = flag.String("host", "localhost", "put your hostname")
	var user *string = flag.String("user", "root", "put your username")
	var pass *string = flag.String("pass", "", "put your password")

	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Pass : ", *pass)

}
