package main

import (
	"crypto/sha1"
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func sliceMap[T any, Y any](xs []T, callback func(item T, index int, array []T) Y) []Y {
	size := len(xs)
	result := make([]Y, size, size)
	for i := 0; i < size; i++ {
		result[i] = callback(xs[i], i, xs)
	}
	return result
}

func sliceMapSimple[T any, Y any](xs []T, callback func(item T) Y) []Y {
	return sliceMap(xs, func(item T, index int, array []T) Y {
		return callback(item)
	})
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

func zStat[T Number](xs []T, mu float64) float64 {
	size := float64(len(xs))
	sizeSqrt := math.Sqrt(size)
	mean := arrMean(xs)
	stddev := arrStddev(xs)
	return (mean - mu) / (stddev / sizeSqrt)
}

func createLinearRegression[T Number, Y Number](xs []T, ys []Y) func(x T) Y {
	size := len(xs)
	xy := make([]float64, size, size)
	for i := 0; i < size; i++ {
		xy[i] = float64(xs[i]) * float64(ys[i])
	}
	Exy := arrSum(xy)
	xy = nil
	x2 := make([]float64, size, size)
	for i := 0; i < size; i++ {
		x2[i] = float64(xs[i] * xs[i])
	}
	Ex2 := arrSum(x2)
	x2 = nil
	Ex := arrSum(xs)
	Ey := arrSum(ys)
	N := float64(size)
	m := (N*Exy - Ex*Ey) / (N*Ex2 - Ex*Ex)
	b := (Ey - m*Ex) / N
	return func(x T) Y {
		return Y(m*float64(x) + b)
	}
}

func getLastPart(word string) int {
	h := sha1.New()
	h.Write([]byte(word))
	bs := h.Sum(nil)
	return int(bs[19])
}

func main() {
	x := []float64{1, 2, 3, 4, 5, 6, 7}
	y := []float64{8, 7, 5, 10, 11, 7, 6}
	lineerRegressionModel := createLinearRegression(x, y)
	fmt.Println(lineerRegressionModel(8))
	x2 := sliceMapSimple(x, func(item float64) int {
		return int(item * item)
	})
	fmt.Println(x2)
	fmt.Println(x)
	//start := time.Now()
	//size := 10
	//shard := make([]int, size, size)
	//for i := 100_000; i < 10_000_000; i++ {
	//	//str := fmt.Sprintf("%d", i)
	//	str := strconv.Itoa(i)
	//	lastPart := getLastPart(str)
	//	mod := lastPart % size
	//	shard[mod]++
	//}
	//fmt.Printf("%v\n", shard)
	////stop := time.Now()
	////elapsed := stop.Sub(start)
	//elapsed := time.Since(start)
	//fmt.Printf("Elapsed time in milliseconds: %d ms\n", elapsed.Milliseconds())
	//fmt.Println("stddev:", arrStddev(shard))
}
