package core

type Game struct {
	State GameState
}

type GameState struct {
	Board   Board
	Turn    Color
	InCheck bool
}

func (g Game) MovePiece(x int, y int, X_new int, Y_new int) {

}
