package main

import (
	"fmt"
	"strconv"
)

/*
  A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

  Find the largest palindrome made from the product of two 3-digit numbers.

  Answer = 906609
*/

func stringReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindromeWord(word string) bool {
	return word == stringReverse(word)
}

func solution() int {
	biggest := -1
	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			prod := i * j
			strProd := strconv.Itoa(prod)
			if isPalindromeWord(strProd) {
				if prod > biggest {
					biggest = prod
				}
			}
		}
	}
	return biggest
}

func main() {
	fmt.Println("Solution: ", solution())
}
