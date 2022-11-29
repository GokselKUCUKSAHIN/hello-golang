package main

import (
	"fmt"
	"math/big"
)

/*
  2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

  What is the sum of the digits of the number 2^1000?


  Answer = 1366
*/

func solution() int {
	sum := 0
	x := new(big.Int)
	x.Exp(big.NewInt(2), big.NewInt(1000), nil)
	for _, item := range []rune(x.String()) {
		value := byte(item) - 48
		sum += int(value)
	}
	return sum
}

func main() {
	fmt.Println("Solution: ", solution())
}
