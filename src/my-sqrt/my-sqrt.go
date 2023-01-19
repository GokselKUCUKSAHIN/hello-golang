package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
	"math/rand"
)

const Epsilon = 1e-12

type Number interface {
	constraints.Integer | constraints.Float
}

func floatEquals(val1 float64, val2 float64) bool {
	return math.Abs(val1-val2) < Epsilon
}

func mySqrt[T Number](number T) float64 {
	numberf := float64(number)
	guess := rand.Float64() * numberf * 0.5
	limit := 0.0
	for !floatEquals(guess, limit) {
		guess = (guess + limit) * 0.5
		limit = numberf / guess
	}
	return guess
}

func main() {
	for i := 1; i < 100; i++ {
		valueMy := mySqrt(i)
		valueMath := math.Sqrt(float64(i))
		fmt.Printf("my: %.12f math: %.12f diff: %.20f\n", valueMy, valueMath, valueMy-valueMath)
	}
}

//         sqrt(X) => guess = X / guess
// 		   sqrt(20) => 5 =? 20 / 5 = 4
//		              guess = (guess + number / guess) / 2
