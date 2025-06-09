package main

import (
	"fmt"
)

func main() {
	// go run . or  go run main.go anonymous.go
	a, _ := fooc()
	_, b := fooc()
	fmt.Println(a, b)
	Iota()
}
