package main

import "fmt"

var lru = make(map[int]int)

func fiboMemo(n int) int {
	if n < 2 {
		return 1
	}
	value, exist := lru[n]
	if exist {
		return value
	}
	result := fiboMemo(n-2) + fiboMemo(n-1)
	lru[n] = result
	return result
}

func main() {
	for i := 0; i < 50; i++ {
		fmt.Println(fiboMemo(i))
	}
}
