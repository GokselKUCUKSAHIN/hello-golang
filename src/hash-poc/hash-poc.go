package main

import (
	"crypto/sha1"
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
	"strconv"
	"time"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func arrSum[T Number](xs []T) float64 {
	var sum float64 = 0
	for i := 0; i < len(xs); i++ {
		sum += float64(xs[i])
	}
	return sum
}

func arrMean[T Number](xs []T) float64 {
	size := float64(len(xs))
	if size == 0 {
		return 0
	}
	return arrSum(xs) / size
}

func arrVariance[T Number](xs []T) float64 {
	size := len(xs)
	mean := arrMean(xs)
	temp := make([]float64, size, size)
	for i := 0; i < size; i++ {
		step := float64(xs[i]) - mean
		temp[i] = step * step
	}
	sum := arrSum(temp)
	return sum / float64(size)
}

func arrStddev[T Number](xs []T) float64 {
	return math.Sqrt(arrVariance(xs))
}

func getLastPart(word string) int {
	h := sha1.New()
	h.Write([]byte(word))
	bs := h.Sum(nil)
	return int(bs[19])
}

func main() {
	start := time.Now()
	size := 10
	shard := make([]int, size, size)
	for i := 100_000; i < 10_000_000; i++ {
		//str := fmt.Sprintf("%d", i)
		str := strconv.Itoa(i)
		lastPart := getLastPart(str)
		mod := lastPart % size
		shard[mod]++
	}
	fmt.Printf("%v\n", shard)
	//stop := time.Now()
	//elapsed := stop.Sub(start)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time in milliseconds: %d ms\n", elapsed.Milliseconds())
	fmt.Println("stddev:", arrStddev(shard))
}
