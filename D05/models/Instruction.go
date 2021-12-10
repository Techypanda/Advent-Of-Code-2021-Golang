package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	X1 int
	X2 int
	Y1 int
	Y2 int
}

func NewInstruction(instruction string) *Instruction {
	initialSplit := strings.Split(instruction, "->")
	if len(initialSplit) != 2 {
		panic(fmt.Sprintf("split %s on -> is not 2, so not format of x1,y1 -> x2,y2", instruction))
	}
	pairOne := strings.Split(strings.TrimSpace(initialSplit[0]), ",")
	pairTwo := strings.Split(strings.TrimSpace(initialSplit[1]), ",")
	if len(pairOne) != 2 {
		panic(fmt.Sprintf("Pair one x1,y1 is not of that format: %s", initialSplit[0]))
	}
	if len(pairTwo) != 2 {
		panic(fmt.Sprintf("Pair two x1,y1 is not of that format: %s", initialSplit[1]))
	}
	x1, err := strconv.Atoi(pairOne[0])
	if err != nil {
		panic("x1 is not a int")
	}
	x2, err := strconv.Atoi(pairTwo[0])
	if err != nil {
		panic("x2 is not a int")
	}
	y1, err := strconv.Atoi(pairOne[1])
	if err != nil {
		panic("y1 is not a int")
	}
	y2, err := strconv.Atoi(pairTwo[1])
	if err != nil {
		panic("y2 is not a int")
	}
	inst := Instruction{x1, x2, y1, y2}
	return &inst
}
