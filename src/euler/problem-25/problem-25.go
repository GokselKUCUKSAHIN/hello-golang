package main

import (
	"fmt"
	"math/big"
)

/*
  The Fibonacci sequence is defined by the recurrence relation:

  Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
  Hence, the first 12 terms will be:

  F1 = 1
  F2 = 1
  F3 = 2
  F4 = 3
  F5 = 5
  F6 = 8
  F7 = 13
  F8 = 21
  F9 = 34
  F10 = 55
  F11 = 89
  F12 = 144
  The 12th term, F12, is the first term to contain three digits.

  What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

  Answer = 4782
*/

var (
	bigTwo = big.NewInt(2)
	memo   = make(map[string]*big.Int)
)

func bigFibo(n *big.Int) *big.Int {
	if n.Cmp(bigTwo) == -1 {
		return big.NewInt(1)
	}
	strN := n.String()
	if cached, exist := memo[strN]; exist {
		return cached
	}
	nTwo := big.NewInt(0).Sub(n, big.NewInt(2))
	nOne := big.NewInt(0).Sub(n, big.NewInt(1))
	result := big.NewInt(0).Add(bigFibo(nTwo), bigFibo(nOne))
	memo[strN] = result
	return result
}

func solution() int {
	for i := 0; ; i++ {
		current := bigFibo(big.NewInt(int64(i)))
		if len(current.String()) >= 1000 {
			return i + 1
		}
	}
}

func main() {
	fmt.Println(solution())
}
