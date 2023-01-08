package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func compareTriplets(a []int32, b []int32) []int32 {
	// Write your code here
	winnerScore := []int32{0, 0}
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			winnerScore[0]++
		} else if a[i] < b[i] {
			winnerScore[1]++
		}
	}
	return winnerScore
}

func diagonalDifference(arr [][]int32) int32 {
	var first, second int32 = 0, 0
	// Write your code here
	row := len(arr)
	// col := len(arr[0])
	for i := 0; i < row; i++ {
		si := row - i - 1
		first += arr[i][i]
		second += arr[i][si]
	}
	dif := first - second
	if dif < 0 {
		dif *= -1
	}
	return dif
}

func plusMinus(arr []int32) {
	// Write your code here
	size := float64(len(arr))
	counts := []float64{0.0, 0.0, 0.0}
	for i := 0; i < int(size); i++ {
		if arr[i] > 0 {
			counts[0]++
		} else if arr[i] == 0 {
			counts[2]++
		} else {
			counts[1]++
		}
	}
	fmt.Printf("%.6f\n", counts[0]/size)
	fmt.Printf("%.6f\n", counts[1]/size)
	fmt.Printf("%.6f\n", counts[2]/size)
}

func staircase(n int32) {
	// Write your code here
	size := int(n)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func miniMaxSum(arr []int32) {
	// Write your code here
	var min, max int64 = 0, 0
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := 0; i < 4; i++ {
		min += int64(arr[i])
	}
	for i := 1; i < 5; i++ {
		max += int64(arr[i])
	}
	fmt.Printf("%d %d", min, max)
}

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	sort.Slice(candles, func(i, j int) bool {
		return candles[i] > candles[j]
	})
	var tallest, count = candles[0], 1
	for i := 1; i < len(candles); i++ {
		if candles[i] == tallest {
			count++
		}
	}
	return int32(count)
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func timeConversion(s string) string {
	// Write your code here
	// 12:01:45 AM => true
	isAM := s[8:] == "AM"
	if hour, err := strconv.ParseInt(s[:2], 10, 32); err != nil {
		if (hour >= 12 && isAM) || (hour < 12 && !isAM) {
			hour = (hour + 12) % 24
		}
		return fmt.Sprintf("%02d%s", hour, s[2:8])
	}
	return "err"
}

func missingNumbers(arr, brr []int32) []int32 {
	// Write your code here
	firstMap := make(map[int32]int32)
	secondMap := make(map[int32]int32)
	for _, v := range arr {
		firstMap[v]++
	}
	for _, v := range brr {
		secondMap[v]++
	}
	result := make([]int32, 0)
	for k, v := range secondMap {
		if v != firstMap[k] {
			result = append(result, k)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	counts := make(map[int32]int)
	size := int(n)
	for i := 0; i < size; i++ {
		counts[ar[i]]++
	}
	numPairs := 0
	for _, count := range counts {
		numPairs += count / 2
	}
	return int32(numPairs)
}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	if (x1 < x2 && v1 < v2) || (x2 < x1 && v2 < v1) {
		return "NO"
	}
	for i := int32(0); i < 100000; i++ {
		k1 := x1 + v1*i
		k2 := x2 + v2*i
		if k1 == k2 {
			return "YES"
		}
	}
	return "NO"
}

func gradeRound(grade int32) int32 {
	if grade < 38 {
		return grade
	}
	for i := int32(1); i < 3; i++ {
		rounded := grade + i
		if rounded%5 == 0 {
			return rounded
		}
	}
	return grade
}

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	size := len(grades)
	result := make([]int32, size)
	for i := 0; i < size; i++ {
		result[i] = gradeRound(grades[i])
	}
	return result
}

func caesarCipher(s string, k int32) string {
	// Write your code here
	size := len(s)
	chars := []rune(s)
	for i := 0; i < size; i++ {
		char := chars[i]
		if char < 'A' || (char > 'Z' && char < 'a') || char > 'z' {
			continue
		}
		// char += k % 26
		if char <= 'Z' {
			// upper case
			char = (char + k%26) - 'A'
		}
		if char > 'z' {
			// lower case
			char = 'a' + (char - 'z') - 1
		} else if char >= 'Z' && char < 'a' {
			// upper case
			char = 'A' + (char - 'Z') - 1
		}
		chars[i] = char
	}
	return string(chars)
}

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	var valleyCount, level, valleyFlag = 0, 0, false
	for _, v := range path {
		if v == 'D' {
			next := level - 1
			if !valleyFlag && next < 0 {
				valleyFlag = true
			}
			level = next
		} else {
			// U
			next := level + 1
			if valleyFlag && next >= 0 {
				valleyFlag = false
				valleyCount++
			}
			level = next
		}
	}
	return int32(valleyCount)
}

func valleyPermutation(count int) []string {
	size := int(math.Pow(2, float64(count)))
	permutations := make([]string, size, size)
	for i := 0; i < size; i++ {
		binary := fmt.Sprintf(fmt.Sprintf("%c0%db", '%', count), i)
		downs := strings.ReplaceAll(binary, "0", "D")
		ups := strings.ReplaceAll(downs, "1", "U")
		permutations[i] = ups
	}
	return permutations
}

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	var sum int32 = 0
	for i := 0; i < len(bill); i++ {
		if int32(i) != k {
			sum += bill[i]
		}
	}
	diff := b - sum/2
	if diff == 0 {
		fmt.Println("Bon Appetit")
		return
	}
	fmt.Println(diff)
}

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	result := []int32{0, 0}
	var min, max = scores[0], scores[0]
	for i := 1; i < len(scores); i++ {
		current := scores[i]
		if current > max {
			max = current
			result[0]++
		}
		if current < min {
			min = current
			result[1]++
		}
	}
	return result
}

func main() {
	fmt.Println(breakingRecords([]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}))
}
