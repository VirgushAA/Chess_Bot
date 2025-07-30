package core

import (
	"fmt"
	"strings"
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

func (b *Board) Initialize() {
	for i := 8; i < 16; i++ {
		b.Board[i] = EncodePiece(Pawn, White)
	}
	for i := 48; i < 56; i++ {
		b.Board[i] = EncodePiece(Pawn, Black)
	}
	b.Board[0] = EncodePiece(Rook, White)
	b.Board[7] = EncodePiece(Rook, White)
	b.Board[56] = EncodePiece(Rook, Black)
	b.Board[63] = EncodePiece(Rook, Black)
	b.Board[1] = EncodePiece(Knight, White)
	b.Board[6] = EncodePiece(Knight, White)
	b.Board[57] = EncodePiece(Knight, Black)
	b.Board[62] = EncodePiece(Knight, Black)
	b.Board[2] = EncodePiece(Bishop, White)
	b.Board[5] = EncodePiece(Bishop, White)
	b.Board[58] = EncodePiece(Knight, Black)
	b.Board[61] = EncodePiece(Knight, Black)
	b.Board[3] = EncodePiece(Queen, White)
	b.Board[4] = EncodePiece(King, White)
	b.Board[59] = EncodePiece(Queen, Black)
	b.Board[60] = EncodePiece(King, Black)
}

func (b *Board) GetPieceType(index int) PieceType {
	value := b.Board[index]
	return PieceType(value & 0x7) // Маска 0000 0111
}

func (b *Board) GetPieceColor(index int) Color {
	value := b.Board[index]
	return Color((value >> 3) & 0x1) // маска 0000 0001
}

func EncodePiece(pieceType PieceType, color Color) uint8 {
	return uint8(pieceType) | uint8(color)<<4
}

func DecodePiece(val uint8) (pt PieceType, color Color) {
	pt = PieceType(val & 0x07)
	color = Color(val >> 4)
	return
}

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
	pieceSymbols := [7]string{".", "P", "N", "B", "R", "Q", "K"}

	for r := 7; r >= 0; r-- { // печатаем с 8 по 1 (чтобы белые были снизу, как в шахматах)
		for c := 0; c < 8; c++ {
			p := b.At(r, c)
			pt, color := DecodePiece(p)
			sym := pieceSymbols[pt]
			if color == Black {
				sym = strings.ToLower(sym)
			}
			fmt.Print(sym, " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

// func (b *Board) Print() {
// 	symbol := func(val int8) string {
// 		if val > 16 {
// 			val = (val - 17) * -1
// 		}
// 		switch val {
// 		case 0:
// 			return "."
// 		case 1:
// 			if val > 0 {
// 				return "P"
// 			} else {
// 				return "p"
// 			}
// 		case 2:
// 			if val > 0 {
// 				return "N"
// 			} else {
// 				return "n"
// 			}
// 		case 3:
// 			if val > 0 {
// 				return "B"
// 			} else {
// 				return "b"
// 			}
// 		case 4:
// 			if val > 0 {
// 				return "R"
// 			} else {
// 				return "r"
// 			}
// 		case 5:
// 			if val > 0 {
// 				return "Q"
// 			} else {
// 				return "q"
// 			}
// 		case 6:
// 			if val > 0 {
// 				return "K"
// 			} else {
// 				return "k"
// 			}
// 		}
// 		return "?"
// 	}

// 	for r := 0; r < 8; r++ {
// 		for c := 0; c < 8; c++ {
// 			fmt.Print(symbol(int8(b.At(r, c))), " ")
// 		}
// 		fmt.Println()
// 	}
// }

// func (b *Board) Print() {
// 	for i := 0; i < 64; i++ {
// 		pt := b.GetPieceType(i)
// 		color := b.GetPieceColor(i)
// 		fmt.Print(pieceToChar(pt, color), " ")
// 		if (i+1)%8 == 0 {
// 			fmt.Println()
// 		}
// 	}
// }

// func pieceToChar(pt PieceType, color Color) string {
// 	chars := [...]string{"·", "P", "N", "B", "R", "Q", "K"}
// 	ch := chars[pt]
// 	if color == Black {
// 		return strings.ToLower(ch)
// 	}
// 	return ch
// }

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
