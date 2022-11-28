package ho_fn

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func GArrMap[T Number](arr []T, fn func(i int, item T, arr []T) T) []T {
	resultArr := make([]T, len(arr))
	for i, item := range arr {
		resultArr[i] = fn(i, item, arr)
	}
	return resultArr
}

func GArrReduce[T Number](arr []T, fn func(cur T, acc T, arr []T) T, start T) T {
	acc := start
	for _, item := range arr {
		acc = fn(item, acc, arr)
	}
	return acc
}

func GArrFilter[T Number](arr []T, fn func(item T, i int, arr []T) bool) []T {
	resultArr := make([]T, 0, len(arr)/2)
	for i, item := range arr {
		if fn(item, i, arr) {
			resultArr = append(resultArr, item)
		}
	}
	return resultArr
}
