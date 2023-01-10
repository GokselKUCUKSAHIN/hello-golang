package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func FindMinimum[Number constraints.Ordered](numberArray []Number) Number {
	size := len(numberArray)
	if size < 1 {
		panic("Empty Array!")
	}
	minimum := numberArray[0]
	for i := 1; i < size; i++ {
		if numberArray[i] < minimum {
			minimum = numberArray[i]
		}
	}
	return minimum
}

func main() {
	xs := []int32{3, 2, 1, 5, 6, 7, 1, 2, 3, 4, -5, 6, 10}
	fmt.Println(FindMinimum(xs))
}
