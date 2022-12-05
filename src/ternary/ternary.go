package main

import (
	"fmt"
)

func Ternary[T any](condition bool, If, Else T) T {
	if condition {
		return If
	}
	return Else
}

func main() {
	t := Ternary[string]
	value := 150
	result := t(value < 0, "Negative", t(value > 0, "Positive", "ZERO"))
	fmt.Println(result)
}
