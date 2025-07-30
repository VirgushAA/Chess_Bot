package main

import (
	. "Chess_Bot/core"
	"fmt"
)

var ()

func main() {
	board := &Board{}
	board.Initialize()
	board.Print()

}

func RenderBoard(state GameState, flipped bool) {
	symbols := make([][]rune, 8)
	for rank := 0; rank < 8; rank++ {
		symbols[rank] = make([]rune, 8)
		for file := 0; file < 8; file++ {
			index := rank*8 + file
			piece, _ := DecodePiece(state.Board[index])
			symbols[rank][file] = pieceToRune(piece)
		}
	}

	// теперь symbols[8][8] — это матрица символов
	printMatrix(symbols, flipped)
}

func printMatrix(symbols [][]rune, flipped bool) {
	if flipped {
		for rank := 0; rank < 8; rank++ {
			fmt.Printf("%d ", rank+1)
			for file := 7; file >= 0; file-- {
				fmt.Printf("%c ", symbols[rank][file])
			}
			fmt.Println()
		}
		fmt.Println("  H G F E D C B A")
	} else {
		for rank := 7; rank >= 0; rank-- {
			fmt.Printf("%d ", rank+1)
			for file := 0; file < 8; file++ {
				fmt.Printf("%c ", symbols[rank][file])
			}
			fmt.Println()
		}
		fmt.Println("  A B C D E F G H")
	}
}

func pieceToRune(val uint8) rune {
	pt, color := DecodePiece(val)
	switch pt {
	case Pawn:
		if color == White {
			return 'P'
		}
		return 'p'
	case Knight:
		if color == White {
			return 'N'
		}
		return 'n'
	// и так далее
	default:
		return '.'
	}
}


// cli/main.go

package main

import (
    "fmt"
    "Chess_Bot/core"
)

func main() {
    game := core.NewGame()

    for {
        renderBoard(game.State(), false) // функция view прямо здесь
        // потом обработка ходов, например:
        fmt.Println("Enter move:")
        // ...
    }
}

func renderBoard(state core.GameState, flipped bool) {
    for rank := 0; rank < 8; rank++ {
        r := rank
        if flipped {
            r = 7 - rank
        }
        fmt.Printf("%d ", r+1)
        for file := 0; file < 8; file++ {
            f := file
            if flipped {
                f = 7 - file
            }
            index := r*8 + f
            pt, color := core.DecodePiece(state.Board[index])
            fmt.Printf("%c ", pieceToRune(pt, color))
        }
        fmt.Println()
    }
    fmt.Println("  A B C D E F G H")
}

func pieceToRune(pt core.PieceType, color core.Color) rune {
    // минимальный пример
    switch pt {
    case core.Pawn:
        if color == core.White {
            return 'P'
        }
        return 'p'
    case core.Knight:
        if color == core.White {
            return 'N'
        }
        return 'n'
    case core.Bishop:
        if color == core.White {
            return 'B'
        }
        return 'b'
    case core.Rook:
        if color == core.White {
            return 'R'
        }
        return 'r'
    case core.Queen:
        if color == core.White {
            return 'Q'
        }
        return 'q'
    case core.King:
        if color == core.White {
            return 'K'
        }
        return 'k'
    default:
        return '.'
    }
}