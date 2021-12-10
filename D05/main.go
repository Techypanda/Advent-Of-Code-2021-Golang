package main

import (
	"fmt"

	"techytechster.com/aoc-2021/d05/controllers"
)

func main() {
	instructions := controllers.LoadInstructions("./static/input.txt")
	seaMap := controllers.ConstructMap(instructions)
	controllers.InvokeInstructions(instructions, seaMap)
	fmt.Println(controllers.DangerCount(2, seaMap))
}
