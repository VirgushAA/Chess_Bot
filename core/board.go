package core

import (
	// "strings"
	"fmt"
)

type Board struct {
	Grid [64]int
}

func NewBoard() *Board {
	b := &Board{Grid: [64]int{-4, -3, -2, -5, -6, -2, -3, -4,
		-1, -1, -1, -1, -1, -1, -1, -1,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		1, 1, 1, 1, 1, 1, 1, 1,
		4, 3, 2, 5, 6, 2, 3, 4,
	},
	}

	return b
}
func (b *Board) At(row, col int) int {
	return b.Grid[row*8+col]
}
func (b *Board) Set(row, col, val int) {
	b.Grid[row*8+col] = val
}
func (b *Board) Print() {
	symbol := func(val int) string {
		switch val {
		case 0:
			return "."
		case 1, -1:
			if val > 0 {
				return "P"
			} else {
				return "p"
			}
		case 2, -2:
			if val > 0 {
				return "N"
			} else {
				return "n"
			}
		case 3, -3:
			if val > 0 {
				return "B"
			} else {
				return "b"
			}
		case 4, -4:
			if val > 0 {
				return "R"
			} else {
				return "r"
			}
		case 5, -5:
			if val > 0 {
				return "Q"
			} else {
				return "q"
			}
		case 6, -6:
			if val > 0 {
				return "K"
			} else {
				return "k"
			}
		}
		return "?"
	}

	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			fmt.Print(symbol(b.At(r, c)), " ")
		}
		fmt.Println()
	}
}

// type Board struct {
// 	Squares [8][8]*Piece
// }

// func NewBoard() *Board {
// 	b := &Board{}

// 	for col := 0; col < 8; col++ {
// 		b.Squares[1][col] = &Piece{Type: Pawn, Color: White}
// 		b.Squares[6][col] = &Piece{Type: Pawn, Color: Black}
// 	}

// 	backRank := []PieceType{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}

// 	for col, piecType := range backRank {
// 		b.Squares[0][col] = &Piece{Type: piecType, Color: White}
// 		b.Squares[7][col] = &Piece{Type: piecType, Color: Black}
// 	}

// 	return b
// }

// func (B *Board) Print() {
// 	for row := 7; row >= 0; row-- {
// 		for col := 0; col < 8; col++ {
// 			piece := B.Squares[row][col]
// 			if piece == nil {
// 				print(". ")
// 				continue
// 			}

// 			symbol := pieceSymbol(piece)
// 			print(symbol + " ")
// 		}
// 		println()
// 	}
// }

// func pieceSymbol(p *Piece) string {
// 	symbols := map[PieceType]string{
// 		Pawn:   "P",
// 		Knight: "N",
// 		Bishop: "B",
// 		Rook:   "R",
// 		Queen:  "Q",
// 		King:   "K",
// 	}

// 	s := symbols[p.Type]
// 	if p.Color == Black {
// 		s = strings.ToLower(s)
// 	}

// 	return s
// }
