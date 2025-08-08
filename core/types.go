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

const (
	PawnValue   = 1
	KnightValue = 3
	BishopValue = 3
	RookValue   = 5
	QueenValue  = 9
)

type CastleRights struct {
	WhiteShortCastleRight bool
	WhiteLongCastleRight  bool
	BlackShortCastleRight bool
	BlackLongCastleRight  bool
}
