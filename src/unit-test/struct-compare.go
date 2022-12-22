package main

import (
	"fmt"
	"reflect"
)

//type People struct {
//	name            string
//	surname         string
//	favoriteNumbers []int
//}

func Equals(x, y any) bool {
	first := reflect.ValueOf(x)
	second := reflect.ValueOf(y)

	if first.Type() != second.Type() {
		return false
	}

	if first.Kind() == reflect.Struct && first.NumField() == second.NumField() {
		for i := 0; i < first.NumField(); i++ {
			firstField := first.Field(i)
			secondField := second.Field(i)
			if !Equals(firstField.Interface(), secondField.Interface()) {
				return false
			}
		}
		return true
	}

	if first.Kind() == reflect.Slice {
		if first.Len() != second.Len() {
			return false
		}
		for i := 0; i < first.Len(); i++ {
			if !Equals(first.Index(i).Interface(), second.Index(i).Interface()) {
				return false
			}
		}
		return true
	}

	return first.Interface() == second.Interface()
}

func main() {
	xs := []byte{1, 3, 2, 4, 5}
	ys := []byte{1, 2, 3, 4, 5}
	z1s := []any{1, 2, 3, xs, 45}
	z2s := []any{1, 2, 3, ys, 45}
	fmt.Println(Equals(z1s, z2s))

	// var p1 *People

	//p1 := People{surname: "Test", name: "Hello", favoriteNumbers: []int{1, 2, 3}}
	//p2 := People{name: "Hello", surname: "Test", favoriteNumbers: []int{1, 2, 3}}

	//p1Reflect := reflect.ValueOf(p1)

	//for i := 0; i < 100; i++ {
	//	for j := 0; j < p1Reflect.NumField(); j++ {
	//		fmt.Printf("%v ", p1Reflect.Field(j))
	//	}
	//	fmt.Println()
	//}

}

/*
	v1 := reflect.ValueOf(1)
	v2 := reflect.ValueOf("hello world")
	v3 := reflect.ValueOf('x')
	v4 := reflect.ValueOf([]byte{'1', '3'})
	v5 := reflect.ValueOf(map[string]byte{"string": 1})
	fmt.Println(v1.Kind())
	fmt.Println(v2.Kind())
	fmt.Println(v3.Kind())
	fmt.Println(v4.Kind())
	fmt.Println(v5.Kind())
*/
