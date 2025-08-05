package core

type Game struct {
	GameState GameState
}

type GameState struct {
	Board            Board
	Turn             Color
	InCheck          bool
	Stalemate        bool
	History          []Move
	GameOver         bool
	EnPassant_target int
}

func (g *Game) PlayATurn(pos string) {

	pos_from, pos_to := convertPositionToIndex(pos)
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
	legalMoves := GenerateMoves(move.FromPosition, &g.GameState)
	im_legal := false
	for _, m := range legalMoves {
		if move == m {
			im_legal = true
			break
		}
	}
	if im_legal {
		g.GameState.Board.SetPiece(move.ToPosition, g.GameState.Board.GetPieceType(move.FromPosition), g.GameState.Board.GetPieceColor(move.FromPosition))
		g.GameState.Board.RemovePiece(move.FromPosition)
		g.GameState.History = append(g.GameState.History, move)
		g.GameState.Turn = (g.GameState.Turn + 1) % 2
	}
	return nil
}

func convertPositionToIndex(pos string) (pos_from, pos_to Position) {
	from := pos[:2]
	to := pos[2:]
	pos_from = Position{Row: int(from[1]) - '1', Col: int(from[0]) - 'a'}
	pos_to = Position{Row: int(to[1]) - '1', Col: int(to[0]) - 'a'}
	return
}
