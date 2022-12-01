package main

import (
	"fmt"
)

//type Number interface {
//	constraints.Integer | constraints.Float
//}

type Person struct {
	name    string
	surname string
	number  int
	friends []Person
}

type Lazy[T any] func() T

type LazyList[T any] struct {
	head Lazy[T]
	tail *LazyList[T]
}

func toListRest[T any](xs ...T) *LazyList[T] {
	return toList(xs)
}

func toList[T any](xs []T) *LazyList[T] {
	if xs == nil || len(xs) == 0 {
		return nil
	}
	return &LazyList[T]{
		head: func() T {
			return xs[0]
		},
		tail: toList(xs[1:]),
	}
}

func printLazy[T any](lazyList *LazyList[T]) {
	fmt.Println(lazyList.head())
	if lazyList.tail != nil {
		printLazy(lazyList.tail)
	}
}

func ac[T any](values ...T) []T {
	return values
}

// Main Function
func main() {
	b := Person{name: "Arda", surname: "Küçükşahin", number: 4567}
	c := Person{name: "Furkan", surname: "Eravşar", number: 7864}
	a := Person{name: "Göksel", surname: "Küçükşahin", number: 1234, friends: ac(b, c)}
	fmt.Println(a)

	myLazyList := toListRest(1, 2, 3, 4, 5)
	fmt.Println(myLazyList)
	printLazy(myLazyList)
}

/*
	f1 := &LazyList[int]{
		head: func() int {
			return 0
		},
		tail: nil,
	}

	f2 := &LazyList[int]{
		head: func() int {
			return 1
		},
		tail: f1,
	}

	f3 := &LazyList[int]{
		head: func() int {
			return 2
		},
		tail: f2,
	}
	fmt.Println(f3)
*/
