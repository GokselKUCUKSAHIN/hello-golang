package main

import (
	"fmt"
	"strings"
)

// var forbittenRunes = []string{"{", "}", "`", "\"", "'", "\\", "/", ":", "(", ")", "?", "+", "*"}

func FilterNoSQLInjection(rawInput string) string {
	// current := rawInput
	for _, forbittenRune := range forbittenRunes {
		rawInput = strings.ReplaceAll(rawInput, forbittenRune, "")
	}
	return rawInput
}

// https://xyz-abc-api-serv.alpha-century.x.com/obs/ctgr?si=2748&c=TR"}},{"wc":{"c.lmd":{"value":"q*"}}},{"tm":{"c.cl":"TR

func main() {
	testStr := "tr-TR\"}},{\"wildcard\":{\"content.lastModifiedBy\":{\"value\":\"quok*\"}}},{\"term\":{\"content.culture\":\"tr-TR"
	fmt.Println("OG:", testStr)
	fmt.Println("FILTER:", FilterNoSQLInjection(testStr))
	fmt.Println("%q:", fmt.Sprintf("%q", testStr))

	fmt.Printf("hello %q world\n", "tr-TR")
	fmt.Printf("hello %s world\n", "tr-TR")
	fmt.Println(CountNoSQLInjection(testStr))
	fmt.Println(HasAnyNoSQLInjectionRune(testStr))
}
