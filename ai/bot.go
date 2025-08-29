package ai

import "Chess_Bot/core"

func BestMove(state *core.GameState, depth int) core.Move {
	moves := core.GenerateAllLegalMovesForColor(state)

	bestScore := -999999
	var bestMove core.Move

	for _, move := range moves {
		next := *state
		next.Board = state.Board.Clone()
		next.Turn = state.Turn
		next.History = append([]core.Move(nil), state.History...)

		g := core.Game{GameState: next}
		g.MakeAMove(move)
		score := minimax(&g.GameState, depth, false)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}
	return bestMove
}
