package core

type Game struct {
	State GameState
}

type GameState struct {
	Board   Board
	Turn    Color
	InCheck bool
}

func (G GameState) MovePiece(x int, y int, X_new int, Y_new int) {

}

// MovePiece(5, 2, 5, 4)

// 8	4 3 2 5 6 2 3 4
// 7	1 1 1 1 1 1 1 1
// 6	0 0 0 0 0 0 0 0
// 5	0 0 0 0 0 0 0 0
// 4	0 0 0 0 0 0 0 0
// 3	0 0 0 0 0 0 0 0
// 2	1 1 1 1 1 1 1 1
// 1	4 3 2 5 6 2 3 4

// 	a b c d e f g h

// 8	0 0 0 0 0 0 0 0
// 7	0 0 0 0 1 0 0 0
// 6	0 0 0 0 0 0 0 0
// 5	0 0 0 0 0 0 0 0
// 4	0 0 0 0 0 0 0 0
// 3	0 0 0 0 0 0 0 0
// 2	0 0 0 0 1 0 0 0
// 1	0 0 0 0 0 0 0 0

// 	a b c d e f g h
