package main

import (
	"fmt"
	"strings"
)

func matrixPrint(matrix [][]int) {
	r := len(matrix)
	c := len(matrix[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Printf("%2d", matrix[i][j])
		}
		fmt.Println()
	}
}

func printPattern(n int) {
	// print top half of the pattern
	size := 2*n - 1 // 2*n -1
	matrix := [][]int{
		{4, 4, 4, 4, 4, 4, 4},
		{4, 4, 4, 4, 4, 4, 4},
		{4, 4, 4, 4, 4, 4, 4},
		{4, 4, 4, 4, 4, 4, 4},
		{4, 4, 4, 4, 4, 4, 4},
		{4, 4, 4, 4, 4, 4, 4},
		{4, 4, 4, 4, 4, 4, 4},
	}
	for level := 1; level < n; level++ {
		for r := level; r < size-level; r++ {
			for i := level; i < n; i++ {
				matrix[r][i]--
			}
			for i := n; i < size-level; i++ {
				matrix[r][i]--
			}
		}
	}
	matrixPrint(matrix)
}

func main() {
	printPattern(4)
	fmt.Println(strings.TrimSuffix("12:45PM", "PM"))
}
