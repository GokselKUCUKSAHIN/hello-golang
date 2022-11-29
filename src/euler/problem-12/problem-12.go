package main

import (
	"fmt"
	"math"
)

/*
  The sequence of triangle numbers is generated by adding the natural numbers. So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28. The first ten terms would be:

  1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

  Let us list the factors of the first seven triangle numbers:

   1: 1
   3: 1,3
   6: 1,2,3,6
  10: 1,2,5,10
  15: 1,3,5,15
  21: 1,3,7,21
  28: 1,2,4,7,14,28
  We can see that 28 is the first triangle number to have over five divisors.

  What is the value of the first triangle number to have over five hundred divisors?

  Answer = 76576500
*/

func getTriangleNumber(limit int) int {
	sum := 0
	for i := 0; i <= limit; i++ {
		sum += i
	}
	return sum
}

func getNumberDivisors(limit int) []int {
	divisors := make([]int, 0, 1)
	for i := 1; i <= limit; i++ {
		if limit%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func getNumberOfDivisors(number int) int {
	limit := int(math.Sqrt(float64(number)))
	divisorCount := 0
	for i := 1; i <= limit; i++ {
		if number%i == 0 {
			divisorCount += 2
		}
	}
	return divisorCount
}

func main() {
	found := false
	for i := 0; !found; i++ {
		triangle := getTriangleNumber(i)
		// triangleDivisors := getNumberDivisors(triangle)
		// lenDivisiros := len(triangleDivisors)
		lenDivisors := getNumberOfDivisors(triangle)
		if lenDivisors >= 500 {
			found = true
			triangleDivisors := getNumberDivisors(triangle)
			fmt.Printf("%d: %d = %v\n", i, triangle, triangleDivisors)
		}
	}
}
