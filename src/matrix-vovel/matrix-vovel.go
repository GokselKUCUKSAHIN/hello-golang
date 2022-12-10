package main

import "fmt"

// a q r s t
// u a k e i
// f f o o o

type char = uint8

const vovels = "aeuioAEUIO"

func printMatrix(matrix []string) {
	for _, row := range matrix {
		for _, col := range []char(row) {
			fmt.Printf("%2c", col)
		}
		fmt.Println()
	}
}

func getSize(arr []string) (int, int) {
	return len(arr), len(arr[0])
}

func isVovel(chr char) bool {
	for _, item := range []char(vovels) {
		if item == chr {
			return true
		}
	}
	return false
}

func solve(matrix []string) string {
	printMatrix(matrix)
	h, w := getSize(matrix)
	for y := 0; y < h-1; y++ {
		for x := 0; x < w-1; x++ {
			if !isVovel(matrix[y][x+1]) {
				continue
			}
			if !isVovel(matrix[y+1][x]) {
				continue
			}
			if !isVovel(matrix[y+1][x+1]) {
				continue
			}
			return fmt.Sprintf("%d-%d", y, x)
		}
	}
	return "not found"
}

func main() {
	matrix := []string{"aqrst", "uakei", "ffooo"}
	solution := solve(matrix)
	fmt.Println(solution)
}
