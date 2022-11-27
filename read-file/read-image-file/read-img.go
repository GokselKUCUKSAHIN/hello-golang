package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
)

func main() {
	catFile, err := os.Open("cat.png")
	if err != nil {
		log.Fatal(err)
	}
	defer catFile.Close()

	cat, err := png.Decode(catFile)
	if err != nil {
		log.Fatal(err)
	}

	size := cat.Bounds().Size()
	var w, h = size.X, size.Y
	for y := 0; y < w; y++ {
		for x := 0; x < h; x++ {
			pixel := cat.At(x, y)

			fmt.Println(pixel)
		}
	}
}
