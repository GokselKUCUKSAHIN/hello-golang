package ho_fn

func ArrMap(arr []int, fn func(i int, item int, arr []int) int) []int {
	resultArr := make([]int, len(arr))
	for i, item := range arr {
		resultArr[i] = fn(i, item, arr)
	}
	return resultArr
}

func ArrReduce(arr []int, fn func(cur int, acc int, arr []int) int, start int) int {
	acc := start
	for _, item := range arr {
		acc = fn(item, acc, arr)
	}
	return acc
}

func ArrFilter(arr []int, fn func(item int, i int, arr []int) bool) []int {
	resultArr := make([]int, len(arr))
	//for i, item := range arr {
	//
	//}
	return resultArr
}
