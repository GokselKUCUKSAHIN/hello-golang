package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
)

func main() {
	fmt.Println("hello, test")
	log.Info("test", 123, "456")
}
