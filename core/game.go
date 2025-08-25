package core

import "slices"

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
	WhiteScore       int
	BlackScore       int
	CastleRights     CastleRights
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
	g.GameState.WhiteScore = 0
	g.GameState.BlackScore = 0
	g.GameState.CastleRights = CastleRights{true, true, true, true}
}

func (g *Game) PassMove(pos_from, pos_to Position) Move {
	move := Move{}
	move.FromPosition = pos_from.Row*8 + pos_from.Col
	move.ToPosition = pos_to.Row*8 + pos_to.Col
	return move
}

func (g *Game) MakeAMove(move Move) error {
	legalMoves := GenerateLegalMoves(move.FromPosition, &g.GameState)
	im_legal := slices.Contains(legalMoves, move)
	if im_legal {
		if g.GameState.Board.GetPieceType(move.FromPosition) == Pawn {
			g.processPawnMove(move)
		}
		if g.GameState.Board.GetPieceType(move.ToPosition) != 0 && g.GameState.Board.GetPieceColor(move.FromPosition) != g.GameState.Board.GetPieceColor(move.ToPosition) {
			g.processCapture(move)
		}
		g.GameState.Board.SetPiece(move.ToPosition, g.GameState.Board.GetPieceType(move.FromPosition), g.GameState.Board.GetPieceColor(move.FromPosition))
		g.GameState.Board.RemovePiece(move.FromPosition)
		g.GameState.History = append(g.GameState.History, move)
		g.GameState.Turn = (g.GameState.Turn + 1) % 2
		g.GameState.InCheck = InCheck(&g.GameState)
		nextMoves := GenerateAllLegalMovesForColor(&g.GameState)
		g.updaterGameOver(nextMoves)
	}
	return nil
}

func (g *Game) updaterGameOver(moves []Move) {
	if len(moves) == 0 {
		if g.GameState.InCheck {
			g.GameState.GameOver = true
		} else {
			g.GameState.Stalemate = true
			g.GameState.GameOver = true
		}
	}
}

func (g *Game) processPawnMove(move Move) {
	step := 8
	if g.GameState.Turn == Black {
		step = -8
	}
	if move.ToPosition == move.FromPosition+(step*2) {
		g.GameState.EnPassant_target = move.FromPosition + step
	}
	if move.ToPosition == g.GameState.EnPassant_target {
		g.GameState.Board.SetPiece(move.ToPosition, g.GameState.Board.GetPieceType(move.ToPosition-step), g.GameState.Board.GetPieceColor(move.ToPosition-step))
		g.GameState.Board.RemovePiece(move.ToPosition - step)
	}
	if move.ToPosition/8 == 7 || move.ToPosition/8 == 0 {
		g.handlePromotion(move.FromPosition)
	}
}

func (g *Game) processCapture(move Move) {
	score := 0
	switch g.GameState.Board.GetPieceType(move.ToPosition) {
	case Pawn:
		score = PawnValue
	case Knight:
		score = KnightValue
	case Bishop:
		score = BishopValue
	case Rook:
		score = RookValue
	case Queen:
		score = QueenValue
	}
	if g.GameState.Board.GetPieceColor(move.FromPosition) == Black {
		g.GameState.BlackScore += score
	} else {
		g.GameState.WhiteScore += score
	}
}

func convertPositionToIndex(pos string) (pos_from, pos_to Position) {
	from := pos[:2]
	to := pos[2:]
	pos_from = Position{Row: int(from[1]) - '1', Col: int(from[0]) - 'a'}
	pos_to = Position{Row: int(to[1]) - '1', Col: int(to[0]) - 'a'}
	return
}

func (g *Game) handlePromotion(pos int) {
	// Как-то организовать ввод выбора фигуры для повышения вне очередности ходов, пока всегда в ферзя
	g.GameState.Board.SetPiece(pos, Queen, g.GameState.Board.GetPieceColor(pos))
}
