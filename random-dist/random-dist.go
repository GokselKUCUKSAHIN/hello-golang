package main

import (
	"fmt"
	"github.com/adam-lavrik/go-imath/ix"
	"math/rand"
)

func randomIntRange(min int, max int) int {
	low := ix.Min(min, max)
	high := ix.Max(min, max)
	return low + rand.Intn(high-low)
}

func main() {
	const LIMIT = 1_000_000
	sum := 0
	for i := 0; i < LIMIT; i++ {
		sum += randomIntRange(1, 100)
	}
	avg := float64(sum) / float64(LIMIT)
	fmt.Printf("Sum: %9d, Limit: %d, Average: %.3f", sum, LIMIT, avg)
}
