package main

import "fmt"

type Rectange struct {
	name         string
	w            float64
	h            float64
	GetPerimeter func() float64
	GetArea      func() float64
}

func NewRect(name string, w float64, h float64) (r *Rectange) {
	// constructor
	r = &Rectange{
		GetPerimeter: func() float64 {
			return 2 * (r.h + r.w)
		},
		GetArea: func() float64 {
			return r.h * r.w
		},
	}
	r.name = name
	r.w = w
	r.h = h
	return
}

func main() {
	rect := NewRect("my rect", 3.14, 6.12)
	fmt.Println(rect.GetPerimeter())
}
