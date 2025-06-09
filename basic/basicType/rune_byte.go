package main

import "fmt"

func TraversalString() {
	s := "proof播客"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()
	for _, v := range s {
		fmt.Printf("%v(%c)", v, v)
	}
	fmt.Println()
}
