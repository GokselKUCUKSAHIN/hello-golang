package main

import (
	"fmt"
	"math"
)

func main() {
	// Define the triangle, pentagonal, and hexagonal number functions
	triangleNumber := func(n int64) int64 {
		return n * (n + 1) / 2
	}
	inversePentagonalNumber := func(x int64) float64 {
		return (math.Sqrt(24*float64(x)+1) + 1) / 6
	}
	inverseHexagonalNumber := func(x int64) float64 {
		return (math.Sqrt(8*float64(x)+1) + 1) / 4
	}

	// Define the function that checks if a number is both pentagonal and hexagonal
	isWholeNumber := func(f float64) bool {
		return math.Abs(f-math.Floor(f)) < 0.000001
	}
	isPentagonalHexagonal := func(n int64) bool {
		// Check if the number is pentagonal
		p := inversePentagonalNumber(n)
		if isWholeNumber(p) {
			// Check if the number is hexagonal
			h := inverseHexagonalNumber(n)
			if isWholeNumber(h) {
				return true
			}
		}
		return false
	}

	// Find the next triangle number that is also pentagonal and hexagonal
	for t := int64(286); ; t++ {
		i := triangleNumber(t)
		if isPentagonalHexagonal(i) {
			fmt.Println(i)
			break
		}
	}
}
