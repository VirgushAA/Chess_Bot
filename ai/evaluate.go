package ai

import (
	"Chess_Bot/core"
)

func evaluate(state *core.GameState) int {
	return state.WhiteScore - state.BlackScore
}
