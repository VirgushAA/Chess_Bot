package core

type Game struct {
	GameState GameState
}

type GameState struct {
	Board     Board
	Turn      Color
	InCheck   bool
	Stalemate bool
	History   []Move
}

func (g *Game) MakeAMove(move Move) error {

	g.GameState.History = append(g.GameState.History, move)
	g.GameState.Turn = (g.GameState.Turn + 1) % 2
	return nil
}
func extractPieceType(value uint8) PieceType {
	return PieceType(value & 0x7) // Маска 0000 0111
}
func extractColor(value uint8) Color {
	return Color((value >> 3) & 0x1) // Сдвиг на 3 бита и маска 0000 0001
}
func setPieceValue(value *uint8, pieceType PieceType, color Color) {
	*value = (uint8(pieceType) | (uint8(color) << 3))
}
