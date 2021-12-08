package main

import (
	"fmt"
	"strconv"
	"strings"

	"techytechster.com/aoc-2021/d03/models"
)

func main() {
	diag := models.NewDiagnostic("./static/input.txt")
	gamma, err := strconv.ParseInt(strings.TrimSpace(diag.ComputeGammaRate()), 2, 64)
	if err != nil {
		panic(err)
	}
	epsilon, err := strconv.ParseInt(strings.TrimSpace(diag.ComputeEpislonRate()), 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Gamma: %d, Epsilon: %d, GxE: %d\n", gamma, epsilon, gamma*epsilon)
	oxygen, err := strconv.ParseInt(strings.TrimSpace(diag.Oxygen()), 2, 64)
	if err != nil {
		panic(err)
	}
	co2, err := strconv.ParseInt(strings.TrimSpace(diag.CO2()), 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Oxygen: %d, CO2: %d, OxC: %d\n", oxygen, co2, oxygen*co2)
}
