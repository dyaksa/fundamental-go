package main

import "fmt"

func logging() {
	fmt.Println("function selesai dipanggil")
}

func addValue(value int) {
	defer logging()
	fmt.Println("memanggil function")
	total := 10 / value
	fmt.Println(total)
}

func main() {
	//defer akan dipanggil setelah func selesai di eksekusi
	addValue(0)
}
