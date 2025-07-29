package core

import "strings"

type Board struct {
	Squares [8][8]*Piece
}

func NewBoard() *Board {
	b := &Board{}

	for col := 0; col < 8; col++ {
		b.Squares[1][col] = &Piece{Type: Pawn, Color: White}
		b.Squares[6][col] = &Piece{Type: Pawn, Color: Black}
	}

	backRank := []PieceType{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}

	for col, piecType := range backRank {
		b.Squares[0][col] = &Piece{Type: piecType, Color: White}
		b.Squares[7][col] = &Piece{Type: piecType, Color: Black}
	}

	return b
}

func (B *Board) Print() {
	for row := 7; row >= 0; row-- {
		for col := 0; col < 8; col++ {
			piece := B.Squares[row][col]
			if piece == nil {
				print(". ")
				continue
			}

			symbol := pieceSymbol(piece)
			print(symbol + " ")
		}
		println()
	}
}

func pieceSymbol(p *Piece) string {
	symbols := map[PieceType]string{
		Pawn:   "P",
		Knight: "N",
		Bishop: "B",
		Rook:   "R",
		Queen:  "Q",
		King:   "K",
	}

	s := symbols[p.Type]
	if p.Color == Black {
		s = strings.ToLower(s)
	}

	return s
}
