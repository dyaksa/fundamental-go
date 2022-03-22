package main

import "fmt"

func main() {
	var persons = map[string]string{
		"name":    "Dyaksa",
		"address": "Kenep RT 02 RW 03 Demak",
	}

	persons["title"] = "Programmer"

	fmt.Println(persons)
	fmt.Println(persons["name"])
	fmt.Println(persons["address"])

	var books = make(map[string]string)
	books["title"] = "Go Programming"
	books["author"] = "Dyaksa"
	fmt.Println(books)
	delete(books, "title")
	fmt.Println(books)
}
