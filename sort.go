package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int16
}

type UserSlice []User

func (user UserSlice) Len() int {
	return len(user)
}

func (user UserSlice) Less(i, j int) bool {
	return user[i].Age < user[j].Age
}

func (user UserSlice) Swap(i, j int) {
	user[i], user[j] = user[j], user[i]
}

func main() {
	users := []User{
		{"Dyaksa", 27},
		{"Widodo", 27},
		{"Ganet", 25},
		{"Coky", 35},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
