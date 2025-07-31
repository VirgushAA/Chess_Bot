package main

import (
	. "Chess_Bot/core"
	"fmt"
	// "fmt"
)

var ()

func main() {
	game := &Game{}
	game.GameStart()
	game.GameState.Board.Print()
	for !game.GameState.GameOver {
		print("Введи свой ход! (В формате \"е2-е4\" пожалуйстаю)\n")
		var input string
		fmt.Scanln(&input)
		pos_from, pos_to, err_input := parseMoves(input)
		if err_input != nil {
			fmt.Println(err_input)
			continue
		}
		row_from, col_from, err_from := convertPositionToIndex(pos_from)
		row_to, col_to, err_to := convertPositionToIndex(pos_to)
		if err_from != nil {
			fmt.Println(err_from)
			continue
		} else if err_to != nil {
			fmt.Println(err_to)
			continue
		}
		game.PlayATurn(row_from, col_from, row_to, col_to)
		game.GameState.Board.Print()
	}
}

func parseMoves(moves string) (from, to string, err error) {
	if len(moves) != 5 || moves[2] != '-' {
		return "", "", fmt.Errorf("Я не понимаю что ты хочешь, введи ход нормально!")
	}
	from = moves[:2]
	to = moves[3:]
	return
}

func convertPositionToIndex(pos string) (row, col int, err error) {
	if len(pos) != 2 || pos[0] < 'a' || pos[0] > 'h' || pos[1] < '1' || pos[1] > '8' {
		return 0, 0, fmt.Errorf("Я не понимаю что ты хочешь, введи ход нормально!")
	}
	col = int(pos[0]) - 'a'
	row = int(pos[1]) - '1'
	return
}
func DumpMoveIndex(pos string) {
	row, col, err := convertPositionToIndex(pos)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Позиция %s соответствует координатам: %d %d\n", pos, row, col)
	}
}
