package main

import "fmt"

type Customer struct {
	Name   string
	Addres string
	Age    int
}

func (customer Customer) sayHai(name string) {
	fmt.Println("Hai, nama saya", customer.Name, "perkenalkan saya", name)
}

func main() {
	var cust Customer
	cust.Name = "John"
	cust.Addres = "Kenep"
	cust.Age = 30

	fmt.Println(cust)
	cust.sayHai("winandi")

	cust2 := Customer{
		Name:   "Dyaksa",
		Addres: "Indonesia",
		Age:    30,
	}

	fmt.Println(cust2)
}
