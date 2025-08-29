package ai

import (
	"Chess_Bot/core"
	"math"
)

func minimax(state *core.GameState, depth int, maximazing bool) int {
	if depth == 0 || state.GameOver {
		return evaluate(state)
	}

	moves := core.GenerateAllLegalMovesForColor(state)
	if len(moves) == 0 {
		return evaluate(state)
	}

	if maximazing {
		best := math.MinInt
		for _, move := range moves {
			next := *state
			next.Board = state.Board.Clone()
			next.Turn = state.Turn
			next.History = append(next.History, state.History...)
			next.History = append([]core.Move(nil), state.History...)

			g := core.Game{GameState: next}
			g.MakeAMove(move)
			score := minimax(&g.GameState, depth-1, false)
			if score > best {
				best = score
			}
		}
		return best
	} else {
		best := math.MaxInt
		for _, move := range moves {
			next := *state
			next.Board = state.Board.Clone()
			next.Turn = state.Turn
			next.History = append([]core.Move(nil), state.History...)

			g := core.Game{GameState: next}
			g.MakeAMove(move)
			score := minimax(&g.GameState, depth-1, true)
			if score < best {
				best = score
			}
		}
		return best
	}
}
