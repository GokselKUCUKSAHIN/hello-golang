package main

import "fmt"

func userGreet() string {
	const emoji = "👋"
	const message = "Hello, " + emoji
	return message
}

func main() {
	const message = "Hello, GO!"
	fmt.Println(userGreet())
	for i, item := range []rune(message) {
		fmt.Printf("%d: %c (keyCode = %d)\n", i, item, item)
	}
}
