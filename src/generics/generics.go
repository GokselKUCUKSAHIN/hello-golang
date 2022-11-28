package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func GAdd[T constraints.Ordered](first T, rest ...T) T {
	sum := first
	if len(rest) != 0 {
		for _, item := range rest {
			sum += item
		}
	}
	return sum
}

func main() {
	fmin := GMin[float64]
	fmt.Println(fmin(3.14, 2.15))
	fmt.Println(GAdd[float64](4, 5, 5, 6, 5, 7, 4.2))
}
