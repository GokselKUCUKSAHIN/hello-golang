package main

import (
	"fmt"
)

/*
  2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

  What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

  Answer = 232792560;
*/

func gcd(a int, b int) int {
	if b > a {
		temp := a
		a = b
		b = temp
	}
	m := a % b
	if m == 0 {
		return b
	}
	return gcd(b, m)
}

//func testIt(number int, limit int) bool {
//	state := true
//	for i := limit; i >= 2 && state; i-- {
//		state = (number % i) == 0
//	}
//	return state
//}

func solution() int {
	num := 2520
	for i := 11; i < 21; i++ {
		num *= i / gcd(i, num)
	}
	return num
}

func main() {
	fmt.Println("Solution: ", solution())
}
