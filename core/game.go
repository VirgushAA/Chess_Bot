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

func (g *Game) Play() {
	g.GameStart()
}

func (g *Game) GameStart() {
	g.NewGame()
}

func (g *Game) NewGame() {
	g.GameState = GameState{}
	g.GameState.Board.Initialize()
}

func (g *Game) MakeAMove(move Move) error {

	g.GameState.History = append(g.GameState.History, move)
	g.GameState.Turn = (g.GameState.Turn + 1) % 2
	return nil
}
