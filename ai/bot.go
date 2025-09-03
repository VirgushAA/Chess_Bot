package ai

import (
	"Chess_Bot/core"
	"fmt"
	"math/rand"
	"time"
)

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
	if bestScore == 0 {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		randomIndex := rand.Intn(len(moves))

		bestMove = moves[randomIndex]
	}
	return bestMove
}

func MakeAMoveAI(g *core.Game) error {
	best_move := BestMove(&g.GameState, 2)
	err := g.MakeAMove(best_move)
	if err != nil {
		return fmt.Errorf(err.Error())
	} else {
		return nil
	}
}
