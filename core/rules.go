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

func GenerateLegalMoves(pos int, GameState *GameState) (legal []Move) {
	pseudoLegal := GenerateAllMoves(pos, GameState)
	legal = make([]Move, 0, len(pseudoLegal))
	for _, move := range pseudoLegal {
		if !LeavesKingInCheck(GameState, move) {
			legal = append(legal, move)
		}
	}
	return
}

func GenerateAllMoves(pos int, GameState *GameState) (moves []Move) {
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
		} else if diagPos == gs.EnPassant_target && gs.Board.GetPieceColor(diagPos+dir) != color {
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
		if canCastleShort(GameState, color) {
			moves = append(moves, Move{FromPosition: pos, ToPosition: pos + 2})
		}
		if canCastleLong(GameState, color) {
			moves = append(moves, Move{FromPosition: pos, ToPosition: pos - 2})
		}
	}
	return
}

func canCastleShort(state *GameState, color Color) bool {
	if color == White {
		if !state.CastleRights.WhiteShortCastleRight {
			return false
		}
		emptySquares := []int{5, 6} // f1, g1
		for _, sq := range emptySquares {
			if state.Board.GetPieceType(sq) != none {
				return false
			}
			if IsSquareAttacked(sq, Black, state) {
				return false
			}
		}
		// King's starting square also can't be attacked
		if IsSquareAttacked(4, Black, state) {
			return false
		}
		return true
	}
	if !state.CastleRights.BlackShortCastleRight {
		return false
	}
	emptySquares := []int{61, 62} // f8, g8
	for _, sq := range emptySquares {
		if state.Board.GetPieceType(sq) != none {
			return false
		}
		if IsSquareAttacked(sq, White, state) {
			return false
		}
	}
	if IsSquareAttacked(60, White, state) {
		return false
	}
	return true
}

func canCastleLong(state *GameState, color Color) bool {
	if color == White {
		if !state.CastleRights.WhiteLongCastleRight {
			return false
		}
		emptySquares := []int{1, 2, 3} // b1, c1, d1
		for _, sq := range emptySquares {
			if state.Board.GetPieceType(sq) != none {
				return false
			}
		}
		// Squares king passes through must not be attacked
		if IsSquareAttacked(4, Black, state) {
			return false
		}
		if IsSquareAttacked(3, Black, state) {
			return false
		}
		if IsSquareAttacked(2, Black, state) {
			return false
		}
		return true
	}
	if !state.CastleRights.BlackLongCastleRight {
		return false
	}
	emptySquares := []int{57, 58, 59} // b8, c8, d8
	for _, sq := range emptySquares {
		if state.Board.GetPieceType(sq) != none {
			return false
		}
	}
	if IsSquareAttacked(60, White, state) {
		return false
	}
	if IsSquareAttacked(59, White, state) {
		return false
	}
	if IsSquareAttacked(58, White, state) {
		return false
	}
	return true
}

// func LeavesKingInCheck(state *GameState, move Move) bool {
// 	captured := MakeMoveOnState(state, move)
// 	inCheck := InCheck(state)
// 	UndoMoveOnState(state, move, captured)
// 	return inCheck
// }

// func MakeMoveOnState(state *GameState, move Move) (capturedPiece uint8) {
// 	capturedPiece = state.Board.Board[move.ToPosition]
// 	state.Board.SetPiece(move.ToPosition,
// 		state.Board.GetPieceType(move.FromPosition),
// 		state.Board.GetPieceColor(move.FromPosition))
// 	state.Board.RemovePiece(move.FromPosition)
// 	return
// }

// func UndoMoveOnState(state *GameState, move Move, capturedPiece uint8) {
// 	state.Board.SetPiece(move.FromPosition,
// 		state.Board.GetPieceType(move.ToPosition),
// 		state.Board.GetPieceColor(move.ToPosition))
// 	state.Board.SetPiece(move.ToPosition, state.Board.GetPieceType(int(capturedPiece)), state.Board.GetPieceColor(int(capturedPiece))) // Set original piece back
// }

func LeavesKingInCheck(GameState *GameState, move Move) bool {
	copyState := *GameState
	copyState.Board = GameState.Board.Clone()

	copyState.Board.SetPiece(move.ToPosition, copyState.Board.GetPieceType(move.FromPosition), copyState.Board.GetPieceColor(move.FromPosition))
	copyState.Board.RemovePiece(move.FromPosition)

	if copyState.Board.GetPieceType(move.FromPosition) == Pawn &&
		(move.ToPosition/8 == 0 || move.ToPosition/8 == 7) {
		copyState.Board.SetPiece(move.ToPosition, Queen, copyState.Turn) // auto queen for now
	}

	return InCheck(&copyState)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func InCheck(GameState *GameState) bool {
	kingPos := -1
	for sq := 0; sq < 64; sq++ {
		if GameState.Board.GetPieceType(sq) == King &&
			GameState.Board.GetPieceColor(sq) == GameState.Turn {
			kingPos = sq
			break
		}
	}
	if kingPos == -1 {
		panic("No king found for current player") // idk
	}

	return IsSquareAttacked(kingPos, (GameState.Turn+1)%2, GameState)
}

func GenerateAllLegalMovesForColor(state *GameState) []Move {
	moves := make([]Move, 0, 64*8)

	for sq := 0; sq < 64; sq++ {
		if state.Board.GetPieceColor(sq) == state.Turn {
			pseudoLegal := GenerateAllMoves(sq, state)
			for _, m := range pseudoLegal {
				if !LeavesKingInCheck(state, m) {
					moves = append(moves, m)
				}
			}
		}
	}
	return moves
}

func IsSquareAttacked(square int, attacker Color, state *GameState) bool {
	// 1. Pawn attacks
	if attacker == White {
		if square%8 != 0 && square-9 >= 0 && state.Board.GetPieceColor(square-9) == White {
			if state.Board.GetPieceType(square-9) == Pawn {
				return true
			}
		}
		if square%8 != 7 && square-7 >= 0 && state.Board.GetPieceColor(square-7) == White {
			if state.Board.GetPieceType(square-7) == Pawn {
				return true
			}
		}
	} else {
		if square%8 != 0 && square+7 <= 63 && state.Board.GetPieceColor(square+7) == Black {
			if state.Board.GetPieceType(square+7) == Pawn {
				return true
			}
		}
		if square%8 != 7 && square+9 <= 63 && state.Board.GetPieceColor(square+9) == Black {
			if state.Board.GetPieceType(square+9) == Pawn {
				return true
			}
		}
	}

	// 2. Knight attacks
	knightOffsets := []int{-17, -15, -10, -6, 6, 10, 15, 17}
	for _, offset := range knightOffsets {
		target := square + offset
		if target < 0 || target > 63 {
			continue
		}
		// Prevent wrap-around horizontally
		if abs((square%8)-(target%8)) > 2 {
			continue
		}
		pt, col := DecodePiece(state.Board.Board[target])
		if pt == Knight && col == attacker {
			return true
		}
	}

	// 3. Sliding pieces
	// Bishop / Queen diagonals
	bishopDirs := []int{-9, -7, 7, 9}
	for _, dir := range bishopDirs {
		pos := square
		for {
			file := pos % 8
			pos += dir
			if pos < 0 || pos > 63 {
				break
			}
			// Stop if wrapped horizontally
			if abs((pos%8)-file) > 1 {
				break
			}
			pt, col := DecodePiece(state.Board.Board[pos])
			if pt != none {
				if col == attacker && (pt == Bishop || pt == Queen) {
					return true
				}
				break
			}
		}
	}
	// Rook / Queen orthogonals
	rookDirs := []int{-8, 8, -1, 1}
	for _, dir := range rookDirs {
		pos := square
		for {
			file := pos % 8
			pos += dir
			if pos < 0 || pos > 63 {
				break
			}
			if abs((pos%8)-file) > 0 && (dir == -1 || dir == 1) {
				break
			}
			pt, col := DecodePiece(state.Board.Board[pos])
			if pt != none {
				if col == attacker && (pt == Rook || pt == Queen) {
					return true
				}
				break
			}
		}
	}

	// 4. King adjacency
	kingOffsets := []int{-9, -8, -7, -1, 1, 7, 8, 9}
	for _, offset := range kingOffsets {
		target := square + offset
		if target < 0 || target > 63 {
			continue
		}
		if abs((square%8)-(target%8)) > 1 {
			continue
		}
		pt, col := DecodePiece(state.Board.Board[target])
		if pt == King && col == attacker {
			return true
		}
	}

	return false
}
