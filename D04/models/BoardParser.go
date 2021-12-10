package models

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BoardParser struct {
	instructions []string
	boards       []Board
}

func (b *BoardParser) WhatsMYSTUFF() {
	fmt.Println(b.instructions)
	fmt.Println(b.boards)
}

func (b *BoardParser) PerformInstructions() (int, float64) {
	for _, instruction := range b.instructions {
		for boardNo, board := range b.boards {
			floatVal, err := strconv.ParseFloat(strings.TrimSpace(instruction), 64)
			if err != nil {
				panic(err)
			}
			if board.CheckWin(floatVal, 5) {
				return boardNo, floatVal
			}
		}
	}
	panic("there is no winner!")
}

func recurseUntilOne(b BoardParser, winVal *float64) (Board, float64) {
	if len(b.boards) > 1 {
		winner, winVal := b.PerformInstructions()
		newBoards := []Board{}
		for idx, board := range b.boards {
			if idx != winner {
				newBoards = append(newBoards, board)
			}
		}
		b.boards = newBoards
		return recurseUntilOne(b, &winVal)
	} else {
		winner, winVal := b.PerformInstructions()
		return b.boards[winner], winVal
	}
}

func (b *BoardParser) WhoWinsLast() (Board, float64) {
	copy := BoardParser{b.instructions, b.boards}
	return recurseUntilOne(copy, nil)
}

func SumUnmarked(board Board) float64 {
	var sum float64
	for _, row := range board.board {
		for _, piece := range row {
			if !piece.Hit {
				sum += piece.Value
			}
		}
	}
	return sum
}

func (b *BoardParser) SumUnmarked(boardNo int) float64 {
	var sum float64
	for _, row := range b.boards[boardNo].board {
		for _, piece := range row {
			if !piece.Hit {
				sum += piece.Value
			}
		}
	}
	return sum
}

func (b *BoardParser) DescribeBoard(num int) {
	for _, row := range b.boards[num].board {
		for _, piece := range row {
			fmt.Print(piece.Value, "(", piece.Hit, ") ")
		}
		fmt.Println()
	}
}

func DescribeBoard(board Board) {
	for _, row := range board.board {
		for _, piece := range row {
			fmt.Print(piece.Value, "(", piece.Hit, ") ")
		}
		fmt.Println()
	}
}

func processABoard(rawText []string, idx int) (*Board, int) {
	currentSlice := rawText[idx:]
	b := [][]float64{}
	for i := 0; i < len(currentSlice); i++ {
		if len(currentSlice[i]) == 1 {
			break
		}
		iter := []float64{}
		for _, v := range strings.Fields(currentSlice[i]) {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				panic(err)
			}
			iter = append(iter, num)
		}
		b = append(b, iter)
	}
	return NewBoard(b), idx + 6
}

func NewBoardParser(fileLoc string) *BoardParser {
	boards := []Board{}
	outBoardP := BoardParser{}
	dat, err := os.ReadFile(fileLoc)
	if err != nil {
		panic(err)
	}
	splitUp := strings.Split(string(dat), "\n")
	outBoardP.instructions = strings.Split(splitUp[0], ",")
	board, idx := processABoard(splitUp, 2)
	boards = append(boards, *board)
	for idx < len(splitUp) {
		board, idx = processABoard(splitUp, idx)
		boards = append(boards, *board)
	}
	outBoardP.boards = boards
	return &outBoardP
}
