package main

import (
	"fmt"
	"math/big"
)

/*
   n! means n × (n − 1) × ... × 3 × 2 × 1

   For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
   and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

   Find the sum of the digits in the number 100!

   Answer = 648
*/

var (
	bigTwo = big.NewInt(2)
)

func bigFactorial(n *big.Int) *big.Int {
	if n.Cmp(bigTwo) == -1 {
		return n
	}
	return big.NewInt(0).Mul(n, bigFactorial(big.NewInt(0).Sub(n, big.NewInt(1))))
}

func solution() int {
	fac100 := bigFactorial(big.NewInt(100))
	sum := 0
	for _, item := range []rune(fac100.String()) {
		value := byte(item) - 48
		sum += int(value)
	}
	return sum
}

func main() {
	fmt.Println(solution())
}
