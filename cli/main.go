package main

import (
	. "Chess_Bot/core"
	"fmt"
	"strconv"
)

var ()

func main() {
	game := &Game{}
	game.GameStart()
	PrintField(game.GameState)
	for !game.GameState.GameOver {
		print("Введи свой ход! (В формате \"е2-е4\" пожалуйста)\n")
		var input string
		fmt.Scanln(&input)
		pos, err_input := processMoves(input)
		if err_input != nil {
			fmt.Println(err_input)
			continue
		}
		game.PlayATurn(pos)
		PrintField(game.GameState)
	}
}

func PrintField(GameState GameState) {
	fmt.Println()
	fmt.Println("\n   _________________________________")
	for row_ := 7; row_ >= 0; row_-- {
		row := row_
		if GameState.Turn == Black {
			row = 7 - row
		}
		fmt.Printf("%d  ", row+1)
		fmt.Print("|")
		for col_ := 0; col_ <= 7; col_++ {
			col := col_
			if GameState.Turn == Black {
				col = 7 - col
			}
			var symbol string
			color := GameState.Board.GetPieceColor(row*8 + col)
			PieceType := GameState.Board.GetPieceType(row*8 + col)
			switch PieceType {
			case 0:
				symbol = "  "
			case Pawn:
				if color == White {
					symbol = " ♟"
				} else {
					symbol = " ♙"
				}
			case Knight:
				if color == White {
					symbol = " ♞"
				} else {
					symbol = " ♘"
				}
			case Bishop:
				if color == White {
					symbol = " ♝"
				} else {
					symbol = " ♗"
				}
			case Rook:
				if color == White {
					symbol = " ♜"
				} else {
					symbol = " ♖"
				}
			case Queen:
				if color == White {
					symbol = " ♛"
				} else {
					symbol = " ♕"
				}
			case King:
				if color == White {
					symbol = " ♚"
				} else {
					symbol = " ♔"
				}
			default:
				panic(fmt.Sprintf("Неизвестная фигура, не знаю, как мы тут оказались %v", PieceType))
			}
			fmt.Print(symbol + " |")

		}
		fmt.Print("  |     ")
		if row == 5 {
			fmt.Printf("Сейчас ходит мистер %s, его счёт: %d", ConvertTurnColor(GameState), ConvertTurnScore_toInt(GameState, GameState.Turn))
		}
		fmt.Println("\n   _________________________________")
		// fmt.Println()
	}
	if GameState.Turn == Black {
		fmt.Println("\n     H   G   F   E   D   C   B   A\n")
	} else {
		fmt.Println("\n     A   B   C   D   E   F   G   H\n")
	}
	printSidebar(GameState)
}

func printSidebar(GameState GameState) {

}

func processMoves(moves string) (processed string, err error) {
	if !((len(moves) == 5 && moves[2] == '-') || len(moves) == 4) {
		return "", fmt.Errorf("Я не понимаю что ты хочешь, введи ход нормально!")
	}
	var from, to string
	if len(moves) == 5 {
		from = moves[:2]
		to = moves[3:]
	} else {
		from = moves[:2]
		to = moves[2:]
	}
	if len(from) != 2 || from[0] < 'a' || from[0] > 'h' || from[1] < '1' || from[1] > '8' {
		return "", fmt.Errorf("Я не понимаю что ты хочешь, введи ход нормально!")
	}
	if len(to) != 2 || to[0] < 'a' || to[0] > 'h' || to[1] < '1' || to[1] > '8' {
		return "", fmt.Errorf("Я не понимаю что ты хочешь, введи ход нормально!")
	}
	processed = from + to
	return
}

func ConvertTurnScore_toInt(GameState GameState, color Color) (score int) {
	if color == White {
		score = GameState.WhiteScore
	} else {
		score = GameState.BlackScore
	}
	return
}

func ConvertTurnScore_toString(GameState GameState, color Color) (score string) {
	if color == White {
		score = strconv.Itoa(GameState.WhiteScore)
	} else {
		score = strconv.Itoa(GameState.BlackScore)
	}
	return
}
