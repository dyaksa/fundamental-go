package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := time.RFC3339
	formt, _ := time.Parse(layout, "2006-01-01")
	fmt.Println(formt)
}
