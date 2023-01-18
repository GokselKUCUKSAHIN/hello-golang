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

// https://stage-product-spda-search-api-service.earth.trendyol.com/spm/categories?supplierId=2748&culture=tr-TR"}},{"wildcard":{"content.lastModifiedBy":{"value":"quok*"}}},{"term":{"content.culture":"tr-TR

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
