package main

import "fmt"

func main() {
	const message = "Hello, GO!"
	for i, item := range []rune(message) {
		fmt.Printf("%d: %c (keyCode = %d)\n", i, item, item)
	}
}
