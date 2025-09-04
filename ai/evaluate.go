package ai

import (
	"Chess_Bot/core"
)

func evaluate(state *core.GameState) int {

	materialBalance := state.BlackScore - state.WhiteScore

	positionalBalance := evaluatePosition(*state)

	controlBalande := evaluateControl(*state)

	developmentBalance := evaluateDevelopment(*state)

	safetyBalance := evaluateSafety(*state)

	return materialBalance + positionalBalance + controlBalande + developmentBalance + safetyBalance
}

func evaluatePosition(state core.GameState) int {
	return 0
}

func evaluateControl(state core.GameState) int {
	return 0
}

func evaluateDevelopment(state core.GameState) int {
	return 0
}

func evaluateSafety(state core.GameState) int {
	return 0
}
