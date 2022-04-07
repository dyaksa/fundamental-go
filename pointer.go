package main

import "fmt"

type Address struct {
	City     string
	Name     string
	Province string
}

func ChangeCountryToIndonesia(address *Address) {
	address.City = "Indonesia"
}

func main() {
	address1 := Address{"Demak", "Dyaksa", "Jawa Tengah"}
	address2 := &address1
	var address3 *Address = &address1

	// ganti field address 2 dengan pointer
	address2.City = "Semarang"

	//ganti varibale dengan * jadi address1 juga berubah ref
	*address2 = Address{"Jakarta", "Dyaksa", "Jawa Tengah"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	alamat1 := new(Address)

	alamat1.City = "Purwokerto"

	fmt.Println(alamat1)

	almt := Address{"Salatiga", "Dyaksa", "Jawa Tengah"}
	ChangeCountryToIndonesia(&almt)
	fmt.Println(almt)

}
