package main

import (
	"fmt"
	"math/rand"
)

func rollDice() int {
	return rand.Intn(6) + 1
}

func main() {
	occurrence := make([]int, 12, 12)
	const LIMIT = 1_000_000 // 500
	for i := 0; i < LIMIT; i++ {
		rollSum := rollDice() + rollDice() - 1
		occurrence[rollSum]++
	}
	fmt.Printf("%v\n", occurrence)
}
