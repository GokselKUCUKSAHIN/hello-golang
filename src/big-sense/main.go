package main

import (
	"fmt"
)

type People struct {
	name    string
	surname string
}

func (test *People) valid() bool {
	return len(test.name) > 0 && len(test.surname) > 0
}

func main() {
	me := &People{name: "Göksel", surname: "Kucuksahin"}
	fmt.Println(me.valid())
}
