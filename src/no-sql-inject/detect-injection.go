package main

var forbittenRunes = []string{".", "{", "}", "[", "]", "(", ")", "|", "?", "+", "*", "$", "^"}
var forbittenString = ".{}[]()|?+*$^"

func CountNoSQLInjection(rawString string) int {
	sum := 0
	for i := 0; i < len(rawString); i++ {
		for j := 0; j < len(forbittenString); j++ {
			if rawString[i] == forbittenString[j] {
				sum++
				break
			}
		}
	}
	return sum
}

func HasAnyNoSQLInjectionRune(rawString string) bool {
	for i := 0; i < len(rawString); i++ {
		for j := 0; j < len(forbittenString); j++ {
			if rawString[i] == forbittenString[j] {
				return true
			}
		}
	}
	return false
}
