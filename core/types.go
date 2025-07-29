package core

type PieceType int
type Color int

const (
	White Color = iota
	Black
)

const (
	none PieceType = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

type Position struct {
	Row int
	Col int
}
