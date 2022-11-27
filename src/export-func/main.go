package main

import (
	mymath "export-func/export"
	"fmt"
)

func main() {
	valAdd := mymath.Add(3, 4)
	valSub := mymath.Sub(7, 3)
	valMult := mymath.Mult(3, 5)
	valDiv := mymath.Div(22, 7)
	fmt.Println(valAdd, valSub, valMult, valDiv)
}
