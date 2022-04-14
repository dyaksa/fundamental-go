package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(10)
	for i := 1; i <= data.Len(); i++ {
		data.Value = "value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(i interface{}) {
		fmt.Println(i)
	})
}
