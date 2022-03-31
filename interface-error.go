package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}

}
func main() {
	a, err := Pembagi(10, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("hasil ", a)
	}
}
