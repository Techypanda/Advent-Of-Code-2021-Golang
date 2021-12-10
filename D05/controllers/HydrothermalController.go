package controllers

import (
	"fmt"
	"os"
	"strings"

	"techytechster.com/aoc-2021/d05/models"
)

func LoadInstructions(location string) []models.Instruction {
	dat, err := os.ReadFile(location)
	if err != nil {
		panic(err)
	}
	rawInstructions := strings.Split(string(dat), "\n")
	outInstructions := []models.Instruction{}
	for _, raw := range rawInstructions {
		outInstructions = append(outInstructions, *models.NewInstruction(raw))
	}
	return outInstructions
}

func InvokeInstructions(instructions []models.Instruction, seaMap [][]int) {
	for _, instruction := range instructions {
		if instruction.X1 == instruction.X2 {
			plotAlongY(instruction.Y1, instruction.Y2, instruction.X1, seaMap)
		} else if instruction.Y1 == instruction.Y2 {
			plotAlongX(instruction.X1, instruction.X2, instruction.Y1, seaMap)
		} else {
			plot45(instruction.Y1, instruction.Y2, instruction.X1, instruction.X2, seaMap)
		}
	}
}

func DisplayMap(seaMap [][]int) {
	for _, row := range seaMap {
		for _, val := range row {
			if val != 0 {
				fmt.Print(val)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func DangerCount(dangerMinimum int, seaMap [][]int) int { // Whats Considered DANGER inclusive
	dangerCount := 0
	for _, row := range seaMap {
		for _, val := range row {
			if val >= dangerMinimum {
				dangerCount++
			}
		}
	}
	return dangerCount
}

func plot45(y1 int, y2 int, x1 int, x2 int, seaMap [][]int) { // TODO: Swap out for a actual math approach
	var xIncrement, yIncrement int
	if x1 < x2 && y1 < y2 { // +1x +1y
		xIncrement = 1
		yIncrement = 1
	} else if x1 > x2 && y1 > y2 { // -1x -1y
		xIncrement = -1
		yIncrement = -1
	} else if x1 < x2 && y1 > y2 { // +1x -1y
		xIncrement = 1
		yIncrement = -1
	} else { // -1x +1y
		xIncrement = -1
		yIncrement = 1
	}
	k := x1
	i := y1
	for k != x2 || i != y2 {
		seaMap[i][k] += 1
		i += yIncrement
		k += xIncrement
	}
	seaMap[i][k] += 1
}

func plotAlongY(y1 int, y2 int, x int, seaMap [][]int) {
	i := y1
	for i != y2 {
		seaMap[i][x] += 1
		if i < y2 {
			i++
		} else {
			i--
		}
	}
	seaMap[i][x] += 1
}

func plotAlongX(x1 int, x2 int, y int, seaMap [][]int) {
	i := x1
	for i != x2 {
		seaMap[y][i] += 1
		if i < x2 {
			i++
		} else {
			i--
		}
	}
	seaMap[y][i] += 1
}

func ConstructMap(instructions []models.Instruction) [][]int {
	myMap := [][]int{}
	maxX := 0
	maxY := 0
	for _, inst := range instructions {
		if inst.X1 > maxX {
			maxX = inst.X1
		}
		if inst.X2 > maxX {
			maxX = inst.X2
		}
		if inst.Y1 > maxY {
			maxY = inst.Y1
		}
		if inst.Y2 > maxY {
			maxY = inst.Y2
		}
	}
	for i := 0; i <= maxX; i++ {
		row := []int{}
		for k := 0; k <= maxY; k++ {
			row = append(row, 0)
		}
		myMap = append(myMap, row)
	}
	return myMap
}
