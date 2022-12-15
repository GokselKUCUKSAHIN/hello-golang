package main

import (
	"fmt"
	"math"
)

type char = uint8

const sep string = "+-------+-------+-------+-------+"

func isPrime(value int) bool {
	if value == 2 {
		return true
	}
	if (value & 0x1) == 0 {
		return false
	}
	limit := int(math.Sqrt(float64(value)))
	for i := 3; i <= limit; i += 2 {
		if (value % i) == 0 {
			return false
		}
	}
	return true
}

func plusMinus(value bool) char {
	if value {
		return '+'
	}
	return '-'
}

func createRow(odd, prime bool) string {
	return fmt.Sprintf("|%4c   |%4c   |%4c   |", plusMinus(odd), plusMinus(!odd), plusMinus(prime))
}

func evalRow(value int) string {
	odd := value&0x1 == 1
	prime := isPrime(value)
	return fmt.Sprintf("|%4d   %s", value, createRow(odd, prime))
}

func main() {
	println(sep)
	println("| VALUE |  ODD  |  EVEN | PRIME |")
	for i := 2; i < 15; i++ {
		println(sep)
		println(evalRow(i))
	}
	println(sep)
}
