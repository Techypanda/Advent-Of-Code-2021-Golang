package main

import (
	"fmt"

	"techytechster.com/aoc-2021/d04/models"
)

func main() {
	bParser := models.NewBoardParser("./static/input.txt")
	winner, winnerVal := bParser.PerformInstructions()
	fmt.Println("--- PART 1 ---")
	fmt.Printf("Winner Board: %d, Winner Value: %f, Sum: %f\n", winner+1, winnerVal, winnerVal*bParser.SumUnmarked(winner))
	fmt.Println("--- PART 2 ---")
	finalBoard, finalBoardWinVal := bParser.WhoWinsLast()
	models.DescribeBoard(finalBoard)
	fmt.Printf("Winner Value: %f, Sum: %f\n", finalBoardWinVal, models.SumUnmarked(finalBoard)*finalBoardWinVal)
}
