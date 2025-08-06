package core

type knightMove struct {
	offset int
	badL   bool
	badLL  bool
	badR   bool
	badRR  bool
}

var knightCandidates = []knightMove{
	{-17, true, false, false, false},
	{-15, false, false, true, false},
	{-10, true, true, false, false},
	{-6, false, false, true, true},
	{6, true, true, false, false},
	{10, false, false, true, true},
	{15, true, false, false, false},
	{17, false, false, true, false},
}

type sliderMove struct {
	offset int
	badL   bool
	badR   bool
}

var bishopCandidates = []sliderMove{
	{-9, true, false},
	{-7, false, true},
	{7, true, false},
	{9, false, true},
}

var rookCandidates = []sliderMove{
	{-8, false, false},
	{-1, true, false},
	{1, false, true},
	{8, false, false},
}

func GenerateMoves(pos int, GameState *GameState) (moves []Move) {
	piece_color := GameState.Board.GetPieceColor(pos)

	if GameState.Board.GetPieceColor(pos) != GameState.Turn {
		return
	}

	switch GameState.Board.GetPieceType(pos) {
	case Pawn:
		moves = generatePawnMoves(pos, GameState, piece_color)
	case Knight:
		moves = generateKnightMoves(pos, GameState, piece_color)
	case Bishop:
		moves = generateBishopMoves(pos, GameState, piece_color)
	case Rook:
		moves = generateRookMoves(pos, GameState, piece_color)
	case Queen:
		moves = generateQueenMoves(pos, GameState, piece_color)
	case King:
		moves = generateKingMoves(pos, GameState, piece_color)
	}
	return
}

func generatePawnMoves(pos int, gs *GameState, color Color) (moves []Move) {
	dir := 8
	startRank := 1

	if color == Black {
		dir = -8
		startRank = 6
	}

	row := pos / 8
	col := pos % 8

	oneStep := pos + dir
	if oneStep >= 0 && oneStep < 64 && gs.Board.Board[oneStep] == 0 {
		moves = append(moves, Move{FromPosition: pos, ToPosition: oneStep})
		if row == startRank {
			twoStep := pos + 2*dir
			if gs.Board.Board[twoStep] == 0 {
				moves = append(moves, Move{FromPosition: pos, ToPosition: twoStep})
			}
		}
	}

	for _, dc := range []int{-1, 1} {
		captureCol := col + dc
		if captureCol < 0 || captureCol > 7 {
			continue
		}
		diagPos := pos + dir + dc
		if diagPos < 0 || diagPos >= 64 {
			continue
		}

		target := gs.Board.Board[diagPos]
		if target != 0 {
			_, targetColor := DecodePiece(target)
			if targetColor != color {
				moves = append(moves, Move{FromPosition: pos, ToPosition: diagPos})
			}
		} else if diagPos == gs.EnPassant_target {
			moves = append(moves, Move{FromPosition: pos, ToPosition: diagPos})
		}
	}

	return
}

func generateKnightMoves(pos int, GameState *GameState, color Color) (moves []Move) {
	for _, c := range knightCandidates {
		new_pos := pos + c.offset
		if new_pos < 0 || new_pos > 63 {
			continue
		}
		file := new_pos % 8
		if (c.badL && file == 1) || (c.badLL && file <= 2) || (c.badR && file == 8) || (c.badRR && file >= 7) {
			continue
		}
		target := GameState.Board.Board[new_pos]
		if target == 0 {
			moves = append(moves, Move{FromPosition: pos, ToPosition: new_pos})
		} else {
			_, targetColor := DecodePiece(target)
			if targetColor != color {
				moves = append(moves, Move{FromPosition: pos, ToPosition: new_pos})
			}
		}
	}
	return
}
func generateBishopMoves(pos int, GameState *GameState, color Color) (moves []Move) {
	for _, direction := range bishopCandidates {
		new_pos := pos
		for {
			new_pos += direction.offset
			if new_pos < 0 || new_pos > 63 {
				break
			}
			file := new_pos % 8
			if (direction.badL && file == 7) || (direction.badR && file == 0) {
				break
			}
			target := GameState.Board.Board[new_pos]
			if target == 0 {
				moves = append(moves, Move{FromPosition: pos, ToPosition: new_pos})
			} else {
				_, targetColor := DecodePiece(target)
				if targetColor != color {
					moves = append(moves, Move{FromPosition: pos, ToPosition: new_pos})
				}
				break
			}
		}
	}
	return
}
func generateRookMoves(pos int, GameState *GameState, color Color) (moves []Move) {
	for _, direction := range rookCandidates {
		new_pos := pos
		for {
			new_pos += direction.offset
			if new_pos < 0 || new_pos > 63 {
				break
			}
			file := new_pos % 8
			if (direction.badL && file == 7) || (direction.badR && file == 0) {
				break
			}
			target := GameState.Board.Board[new_pos]
			if target == 0 {
				moves = append(moves, Move{FromPosition: pos, ToPosition: new_pos})
			} else {
				_, targetColor := DecodePiece(target)
				if targetColor != color {
					moves = append(moves, Move{FromPosition: pos, ToPosition: new_pos})
				}
				break
			}
		}
	}
	return
}
func generateQueenMoves(pos int, GameState *GameState, color Color) (moves []Move) {
	moves = append(moves, generateBishopMoves(pos, GameState, color)...)
	moves = append(moves, generateRookMoves(pos, GameState, color)...)
	return
}
func generateKingMoves(pos int, GameState *GameState, color Color) (moves []Move) {
	for _, direction := range append(bishopCandidates, rookCandidates...) {
		new_pos := pos + direction.offset
		if new_pos < 0 || new_pos > 63 {
			continue
		}
		file := new_pos % 8
		if (direction.badL && file == 0) || (direction.badR && file == 7) {
			continue
		}
		target := GameState.Board.Board[new_pos]
		if target == 0 {
			moves = append(moves, Move{FromPosition: pos, ToPosition: new_pos})
		} else {
			_, targetColor := DecodePiece(target)
			if targetColor != color {
				moves = append(moves, Move{FromPosition: pos, ToPosition: new_pos})
			}
		}
	}
	return
}
