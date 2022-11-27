package main

import "fmt"

func main() {
	fmt.Println("Hello")
	defer fmt.Println("Word")
	fmt.Println("Göksel")
	fmt.Println("Go")
}
