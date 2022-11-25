package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	target := "./read-file"
	fileName := "names.txt"
	dat, err := os.ReadDir(target)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("'%s' dir:\n", target)
	for _, dir := range dat {
		//fmt.Println(index, dir)
		fmt.Println(dir.Name())
	}
	fmt.Printf("'%s/%s' content:\n", target, fileName)
	fileContent, err := os.ReadFile(filepath.Join(target, fileName))
	if err != nil {
		fmt.Println(err.Error())
	}
	fileStr := string(fileContent)
	for index, item := range strings.Split(fileStr, "\n") {
		fmt.Println(index, item)
	}
}
