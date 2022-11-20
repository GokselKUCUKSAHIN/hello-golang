package main

import (
	"fmt"
	"math"
)

func isPrime(value int) bool {
	if (value & 0x1) == 0 {
		return false
	}
	limit := int(math.Sqrt(float64(value)))
	for i := 3; i <= limit; i += 2 {
		if (value % i) == 0 {
			return false
		}
	}
	return true
}

func main() {
	const row = 10
	const col = 5

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			current := i*col + j
			if isPrime(current) {
				current = 0
			}
			fmt.Printf("%4d", current)
		}
		fmt.Println()
	}
}
