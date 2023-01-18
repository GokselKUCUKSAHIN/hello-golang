package main

import "fmt"

type MathHandler func(int) int

type MyHandler struct {
	t  MathHandler
	t2 MathHandler
}

func NewMyHandler() *MyHandler {
	return &MyHandler{
		t:  square,
		t2: cube,
	}
}

func square(i int) int {
	return i * i
}

func cube(i int) int {
	return i * i * i
}

////

type Handler struct {
	t  test
	t2 test2
}

type test interface {
	handle(int) int
}

type test2 interface {
	handle(int) int
}

type international struct {
	value int
}

type tr struct {
	value int
}

func (n *international) handle(i int) int {
	return i * i
}

func (n *tr) handle(i int) int {
	return i * i * i
}

func main() {
	//i := international{}
	//t := tr{}

	// single abstract method

	handler := Handler{}

	handler.t = &international{}
	handler.t2 = &tr{}

	fmt.Println(handler.t.handle(5))
	fmt.Println(handler.t2.handle(5))

	// callback

	//myHandler := NewMyHandler()

	myHandler := MyHandler{}

	myHandler.t = square
	myHandler.t2 = cube

	fmt.Println(myHandler.t(5))
	fmt.Println(myHandler.t2(5))
}
