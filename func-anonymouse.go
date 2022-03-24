package main

import "fmt"

type CheckName func(string) bool
type CheckPassword func(string) bool

func registerUser(name string, password string, cName CheckName, cPassword CheckPassword) string {
	if !cName(name) {
		return "username not valid"
	}

	if !cPassword(password) {
		return "password not valid"
	}
	return "user registered"
}

func main() {
	checkPassword := func(password string) bool {
		if password == "secret" {
			return true
		}
		return false
	}

	checkName := func(name string) bool {
		if name == "admin" {
			return true
		}
		return false
	}

	res := registerUser("admin", "halo", checkName, checkPassword)
	fmt.Println(res)
}
