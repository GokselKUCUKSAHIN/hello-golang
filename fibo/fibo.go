package main

import "fmt"

func fibo(n int) int {
	if n < 2 {
		return 1
	}
	return fibo(n-2) + fibo(n-1)
}

func main() {
	for i := 0; i < 50; i++ {
		fmt.Println(fibo(i))
	}
}
