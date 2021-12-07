package main

import (
	"fmt"

	"techytechster.com/aoc-2021/d02/models"
)

func main() {
	sub := models.Submarine{}
	fReader := models.NewFilerReader("./static/input.txt")
	defer fReader.Close() // Garbage Cleanup
	fmt.Println("---- Part One ----")
	line := fReader.NextLine()
	for line != nil {
		sub.InterpretInstruction(*line, 2) // Set to 1 if you want part 1
		line = fReader.NextLine()
	}
	fmt.Printf("Submarine depth: %f\nSubmarine horizontal: %f\nSubmarine depth x horizontal: %f\n", sub.Depth, sub.Horizontal, sub.Depth*sub.Horizontal)
}
