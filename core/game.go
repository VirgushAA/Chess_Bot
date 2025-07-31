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
	GameOver  bool
}

func (g *Game) PlayATurn(row_from, col_from, row_to, col_to int) {
	pos_from := Position{Row: row_from, Col: col_from}
	pos_to := Position{Row: row_to, Col: col_to}

	g.MakeAMove(g.PassMove(pos_from, pos_to))
}

func (g *Game) GameStart() {
	g.NewGame()

}

func (g *Game) NewGame() {
	g.GameState.Board.Initialize()
	g.GameState.Turn = White
	g.GameState.InCheck = false
	g.GameState.Stalemate = false
	g.GameState.GameOver = false
}
func (g *Game) PassMove(pos_from, pos_to Position) Move {
	move := Move{}
	move.FromPosition = pos_from.Row*8 + pos_from.Col
	move.ToPosition = pos_to.Row*8 + pos_to.Col
	return move
}

func (g *Game) MakeAMove(move Move) error {
	g.GameState.Board.SetPiece(move.ToPosition, g.GameState.Board.GetPieceType(move.FromPosition), g.GameState.Board.GetPieceColor(move.FromPosition))
	g.GameState.Board.RemovePiece(move.FromPosition)
	g.GameState.History = append(g.GameState.History, move)
	g.GameState.Turn = (g.GameState.Turn + 1) % 2
	return nil
}
