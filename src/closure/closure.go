package main

import fmt "fmt"

func greeter(greetMessage string) func(name string) string {
	return func(name string) string {
		return fmt.Sprintf("%s, %s", greetMessage, name)
	}
}

func createCounter(initialValue int) func() int {
	counter := initialValue
	return func() int {
		counter++
		return counter
	}
}

func main() {
	hello := greeter("Hello")
	hi := greeter("Hi")
	fmt.Println(hello("Göksel"))
	fmt.Println(hello("Arda"))
	fmt.Println(hi("Göksel"))

	myCounter := createCounter(10)
	fmt.Println(myCounter())
	fmt.Println(myCounter())
	fmt.Println(myCounter())
	fmt.Println(myCounter())
}
