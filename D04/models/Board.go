package models

type BoardPiece struct {
	Value float64
	Hit   bool
}
type Board struct {
	board [][]BoardPiece
}

func (b *Board) CheckWin(val float64, winCondition float64) bool {
	for y, row := range b.board {
		for x, piece := range row {
			if piece.Value == val {
				piece.Hit = true
				b.board[y][x] = piece
				hits := 0
				for i := 0; i < len(row); i++ {
					if b.board[y][i].Hit {
						hits += 1
					}
				}
				if hits == int(winCondition) {
					return true
				}
				hits = 0
				for i := 0; i < len(b.board); i++ {
					if b.board[i][x].Hit {
						hits += 1
					}
				}
				if hits == int(winCondition) {
					return true
				}
			}
		}
	}
	return false
}

func NewBoard(boardVal [][]float64) *Board {
	board := make([][]BoardPiece, len(boardVal))
	for i, row := range boardVal {
		for k, val := range row {
			if board[i] == nil {
				board[i] = make([]BoardPiece, len(row))
			}
			board[i][k] = BoardPiece{val, false}
		}
	}
	return &Board{board}
}
