package main

import (
	"fmt"
	_ "strconv"
)

var connection string

func init() {
	connection = "MYSQL"
}

func getDatabase() string {
	return connection
}

func main() {
	db := getDatabase()
	fmt.Println(db)
	// strconv.AppendInt(nil, 10, 10)
}
