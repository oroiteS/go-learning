package main

import (
	"fmt"
	"math/rand/v2"
)

func sum(arr []int) int64 {
	sum := int64(0)
	for _, v := range arr {
		sum += int64(v)
	}
	return sum
}
func main() {
	var arr1 [1000]int
	for i := 0; i < 1000; i++ {
		arr1[i] = rand.IntN(1000)
	}
	fmt.Println(sum(arr1[:]))
	arr2 := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr2[i] = i
	}
	fmt.Println(sum(arr2[:]))
}
