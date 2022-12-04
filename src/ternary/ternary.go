package main

import "fmt"

func Ternary[T any](condition bool, If, Else T) T {
	if condition {
		return If
	}
	return Else
}

func main() {
	value := 0
	result := Ternary[string](value < 0, "Negative", Ternary[string](value > 0, "Positive", "ZERO"))
	fmt.Println(result)
}
