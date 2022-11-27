package main

import "fmt"

/*
  The prime factors of 13195 are 5, 7, 13 and 29.

  What is the largest prime factor of the number 600851475143?

  Answer = 6857
*/

func solution() int {
	var curr, div, largestDiv = 600851475143, 2, -1
	for ok := true; ok; ok = curr != 1 {
		if curr%div == 0 {
			curr /= div
			if div > largestDiv {
				largestDiv = div
			}
			div = 2
		} else {
			if div == 2 {
				div = 3
			} else {
				div += 2
			}
		}
	}
	return largestDiv
}

func main() {
	fmt.Println("Solution: ", solution())
}
