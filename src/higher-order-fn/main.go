package main

import (
	"fmt"
	ho_fn "higher-oder-fn/ho-fn"
	"math"
)

func main() {
	myArr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(myArr)
	fmt.Printf("len: %d, cap: %d\n", len(myArr), cap(myArr))

	tripler := func(i, item int, arr []int) int {
		return item * 3
	}

	otherArr := ho_fn.ArrMap(myArr, func(i int, item int, arr []int) int {
		return item * 2
	})
	fmt.Println(otherArr)

	tripleArr := ho_fn.ArrMap(myArr, tripler)
	fmt.Println(tripleArr)

	sum := ho_fn.ArrReduce(myArr, func(cur int, acc int, arr []int) int {
		return acc + cur
	}, 0)
	fmt.Println(sum)

	filter := ho_fn.ArrFilter(otherArr, func(item int, i int, arr []int) bool {
		return item > 5
	})
	fmt.Println(filter)

	myFloatArr := []float64{3.14, 15.92, 65.35}
	fmt.Println(myFloatArr)

	sqrtArr := ho_fn.GArrMap(myFloatArr, func(i int, item float64, arr []float64) float64 {
		return math.Sqrt(item)
	})
	fmt.Println(sqrtArr)

}
