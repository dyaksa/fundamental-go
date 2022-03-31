package main

import "fmt"

func endApp() {
	message := recover()
	fmt.Println("error", message)
	fmt.Println("aplikasi selesai di running")
}

func startApp(err bool) {
	defer endApp()
	if !err {
		panic("something wrong")
	}
	fmt.Println("aplikasi berjalan")
}

func main() {
	startApp(false)
}
