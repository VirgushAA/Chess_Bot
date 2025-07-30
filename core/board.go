package core

import (
	// "strings"
	"fmt"
)

type Board struct {
	Board [64]uint8
}

func NewBoard() *Board {
	b := &Board{Board: [64]uint8{4, 3, 2, 5, 6, 2, 3, 4,
		1, 1, 1, 1, 1, 1, 1, 1,
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

// Извлечь тип фигуры
func (b *Board) GetPieceType(index int) PieceType {
	value := b.Board[index]
	return PieceType(value & 0x7) // Маска 0000 0111
}

// Извлечь цвет фигуры
func (b *Board) GetPieceColor(index int) Color {
	value := b.Board[index]
	return Color((value >> 3) & 0x1) // Сдвиг на 3 и маска 0000 0001
}

// Установить тип и цвет фигуры
func (b *Board) SetPiece(index int, pieceType PieceType, color Color) {
	var packedValue uint8 = (uint8(pieceType) | (uint8(color) << 3))
	b.Board[index] = packedValue
}

func (b *Board) At(row, col int) uint8 {
	return b.Board[row*8+col]
}
func (b *Board) Set(row, col int, val uint8) {
	b.Board[row*8+col] = val
}
func (b *Board) Print() {
	symbol := func(val uint8) string {
		switch val {
		case 0:
			return "."
		case 1:
			if val > 0 {
				return "P"
			} else {
				return "p"
			}
		case 2:
			if val > 0 {
				return "N"
			} else {
				return "n"
			}
		case 3:
			if val > 0 {
				return "B"
			} else {
				return "b"
			}
		case 4:
			if val > 0 {
				return "R"
			} else {
				return "r"
			}
		case 5:
			if val > 0 {
				return "Q"
			} else {
				return "q"
			}
		case 6:
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
