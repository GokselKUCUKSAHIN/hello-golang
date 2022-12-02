package main

import (
	"fmt"
)

type Lazy[T any] func() T

type LazyList[T any] struct {
	head Lazy[T]
	tail *LazyList[T]
}

func toListRest[T any](x0 T, xs ...T) *LazyList[T] {
	return toList(append([]T{x0}, xs...))
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

func getLazy[T any](lazyList *LazyList[T], index int) Lazy[T] {
	if index < 0 {
		return nil
	}
	if index == 0 {
		return lazyList.head
	}
	return getLazy(lazyList.tail, index-1)
}

func lenLazy[T any](lazyList *LazyList[T], count int) int {
	if lazyList.tail == nil {
		return count + 1
	}
	return lenLazy(lazyList.tail, count+1)
}

//// not working well
//func lazyRange(begin int) *LazyList[int] {
//	return &LazyList[int]{
//		head: func() int { return begin },
//		tail: lazyRange(begin + 1),
//	}
//}
//
//func take[T any](lazyList *LazyList[T], n int) *LazyList[T] {
//	fmt.Println(lazyList.head())
//	if lazyList.tail != nil && n >= 0 {
//		return take(lazyList.tail, n-1)
//	}
//	// end of list
//	return nil
//}

func ac[T any](values ...T) []T {
	return values
}

// Main Function
func main() {

	myLazyList := toListRest(0, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(myLazyList)
	fmt.Println(lenLazy(myLazyList, 0))
	// printLazy(myLazyList)
	// fmt.Println(getLazy(myLazyList, 2)())
	// printLazy(toListRest(1, 2, 3, 4, 5))
	// printLazy(lazyRange(1))
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
