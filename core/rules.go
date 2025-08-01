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
	{7, false, true},
	{9, true, false},
}

var rookCandidates = []sliderMove{
	{-8, false, false},
	{-1, true, false},
	{1, false, true},
	{8, true, false},
}

func GenerateMoves(pos int, board Board) (moves []Move) {
	piece_color := board.GetPieceColor(pos)

	switch board.GetPieceType(pos) {
	case Pawn:
		moves = generatePawnMoves(pos, board, piece_color)
	case Knight:
		moves = generateKnightMoves(pos, board, piece_color)
	case Bishop:
		moves = generateBishopMoves(pos, board, piece_color)
	case Rook:
		moves = generateRookMoves(pos, board, piece_color)
	case Queen:
		moves = generateQueenMoves(pos, board, piece_color)
	case King:
		moves = generateKingMoves(pos, board, piece_color)
	}
	return
}

func generatePawnMoves(pos int, board Board, color Color) (moves []Move) {
	return
}
func generateKnightMoves(pos int, board Board, color Color) (moves []Move) {
	for _, c := range knightCandidates {
		new_pos := pos + c.offset
		if new_pos < 0 || new_pos > 63 {
			continue
		}
		file := new_pos % 8
		if (c.badL && file == 0) || (c.badLL && file <= 1) || (c.badR && file == 7) || (c.badRR && file >= 6) {
			continue
		}
		target := board.Board[new_pos]
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
func generateBishopMoves(pos int, board Board, color Color) (moves []Move) {
	for _, direction := range bishopCandidates {
		new_pos := pos
		for {
			new_pos += direction.offset
			if new_pos < 0 || new_pos > 63 {
				break
			}
			file := new_pos % 8
			if (direction.badL && file == 0) || (direction.badR && file == 7) {
				break
			}
			target := board.Board[new_pos]
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
func generateRookMoves(pos int, board Board, color Color) (moves []Move) {
	for _, direction := range rookCandidates {
		new_pos := pos
		for {
			new_pos += direction.offset
			if new_pos < 0 || new_pos > 63 {
				break
			}
			file := new_pos % 8
			if (direction.badL && file == 0) || (direction.badR && file == 7) {
				break
			}
			target := board.Board[new_pos]
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
func generateQueenMoves(pos int, board Board, color Color) (moves []Move) {
	moves = append(moves, generateBishopMoves(pos, board, color)...)
	moves = append(moves, generateRookMoves(pos, board, color)...)
	return
}
func generateKingMoves(pos int, board Board, color Color) (moves []Move) {
	for _, direction := range append(bishopCandidates, rookCandidates...) {
		new_pos := pos + direction.offset
		if new_pos < 0 || new_pos > 63 {
			continue
		}
		file := new_pos % 8
		if (direction.badL && file == 0) || (direction.badR && file == 7) {
			continue
		}
		target := board.Board[new_pos]
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
