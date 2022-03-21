package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	fmt.Println(months)

	var slice1 = months[4:8]
	fmt.Println(slice1)
	fmt.Println(slice1[0])

	var slice2 = append(slice1, "Bulan Baru")
	fmt.Println(slice2)
	fmt.Println(slice1)
	fmt.Println(months)

	slice1[0] = "Ganti Nama Bulan"
	fmt.Println(slice1)
	fmt.Println(months)

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	newSlice := make([]string, 7, 7)
	newSlice[0] = "Dyaksa"
	newSlice[1] = "Jaharuddin"
	newSlice[2] = "dan"
	newSlice[3] = "Tifani"
	newSlice[4] = "Maulizha"
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//important hati hati dalam membuat array
	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
