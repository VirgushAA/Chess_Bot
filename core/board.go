package core

import (
	"fmt"
	"strings"
)

type Board struct {
	Board [64]uint8
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

func (b *Board) RemovePiece(index int) {
	b.Board[index] = 0
}

func (b *Board) At(row, col int) uint8 {
	return b.Board[row*8+col]
}

func (b *Board) Set(row, col int, val uint8) {
	b.Board[row*8+col] = val
}

func (b *Board) Print() {
	pieceSymbols := [7]string{".", "P", "N", "B", "R", "Q", "K"}

	for r := 7; r >= 0; r-- { // печатаем с 8 рядв по 1 (чтобы белые были снизу)
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
