package main

import (
	ho_fn "array-map/ho-fn"
	"fmt"
)

func main() {
	myArr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(myArr)
	fmt.Println(len(myArr))
	otherArr := ho_fn.ArrMap(myArr, func(i int, item int, arr []int) int {
		return item * 2
	})
	fmt.Println(otherArr)
	sum := ho_fn.ArrReduce(myArr, func(cur int, acc int, arr []int) int {
		return acc + cur
	}, 0)
	fmt.Println(sum)
	filter := ho_fn.ArrFilter(otherArr, func(item int, i int, arr []int) bool {
		return item > 5
	})
	fmt.Println(filter)
}
