package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Submarine struct {
	Horizontal float64
	Depth      float64
	Aim        float64
}

func (s *Submarine) MoveHorizontal(units float64) {
	s.Horizontal += units
}
func (s *Submarine) MoveAim(units float64) {
	s.Aim += units
}
func (s *Submarine) MoveDepth(units float64) {
	s.Depth += units
}
func (s *Submarine) InterpretInstruction(raw string, part int) {
	rawSplit := strings.Split(raw, " ")
	operation := strings.ToUpper(rawSplit[0])
	units, err := strconv.ParseFloat(rawSplit[1], 64)
	if err != nil {
		panic(err)
	}
	if part == 1 { // Part One's Understanding
		switch operation {
		case "FORWARD":
			s.MoveHorizontal(units)
		case "DOWN":
			s.MoveDepth(units)
		case "UP":
			s.MoveDepth(units * -1)
		default:
			panic(fmt.Sprintf("Undefined Operation - %s", operation))
		}
	} else { // Part Two's Understanding
		switch operation {
		case "FORWARD":
			s.MoveHorizontal(units)
			s.MoveDepth(s.Aim * units)
		case "DOWN":
			s.MoveAim(units)
		case "UP":
			s.MoveAim(units * -1)
		default:
			panic(fmt.Sprintf("Undefined Operation - %s", operation))
		}
	}
}
