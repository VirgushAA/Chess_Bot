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
	g.MakeAMove(g.PassMove(row_from, col_from, row_to, col_to))
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
func (g *Game) PassMove(row_from, col_from, row_to, col_to int) Move {
	move := Move{}
	move.FromPosition = row_from*8 + col_from
	move.ToPosition = row_to*8 + col_to
	return move
}

func (g *Game) MakeAMove(move Move) error {
	g.GameState.Board.SetPiece(move.ToPosition, g.GameState.Board.GetPieceType(move.FromPosition), g.GameState.Board.GetPieceColor(move.FromPosition))
	g.GameState.Board.RemovePiece(move.FromPosition)
	g.GameState.History = append(g.GameState.History, move)
	g.GameState.Turn = (g.GameState.Turn + 1) % 2
	return nil
}
