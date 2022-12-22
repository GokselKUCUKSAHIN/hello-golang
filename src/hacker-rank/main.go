package main

import (
	"fmt"
	"sort"
	"strconv"
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
	isAM := s[8:] == "AM"
	hour, _ := strconv.ParseInt(s[:2], 10, 32)
	if (hour >= 12 && isAM) || (hour < 12 && !isAM) {
		hour = (hour + 12) % 24
	}
	return fmt.Sprintf("%02d%s", hour, s[2:8])
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

	// create a slice to store the missing numbers
	result := make([]int32, 0)

	// iterate over the elements in the second list and check if the frequency is the same in both lists
	for k, v := range secondMap {
		if v != firstMap[k] {
			// frequency is not the same, so add the number to the slice
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
	// Iterate through the array and increment the count for each number
	for i := 0; i < size; i++ {
		counts[ar[i]]++
	}
	// Initialize a counter for the number of pairs
	numPairs := 0
	// Iterate through the counts in the map and add the number of pairs for each count
	for _, count := range counts {
		numPairs += count / 2
	}
	return int32(numPairs)
}

func main() {
	fmt.Println(compareTriplets([]int32{17, 28, 30}, []int32{99, 16, 8}))
	staircase(6)
	fmt.Println(timeConversion("12:01:45PM"))
	fmt.Println(timeConversion("12:01:45AM"))
	fmt.Println(timeConversion("07:05:45PM"))

	// var arr, brr = []int32{7, 2, 5, 3, 5, 3}, []int32{7, 2, 5, 4, 6, 3, 5, 3}
	var arr, brr = []int32{203, 204, 205, 206, 207, 208, 203, 204, 205, 206},
		[]int32{203, 204, 204, 205, 206, 207, 205, 208, 203, 206, 205, 206, 204}

	fmt.Println(missingNumbers(arr, brr))

	socks := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println(sockMerchant(int32(len(socks)), socks))
}
