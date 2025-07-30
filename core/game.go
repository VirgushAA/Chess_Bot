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
