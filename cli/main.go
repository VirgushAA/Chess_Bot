package main

import (
	. "Chess_Bot/core"
	"fmt"
)

var ()

func main() {
	game := &Game{}
	game.GameStart()
	game.GameState.Board.Print()
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
		game.GameState.Board.Print()
	}
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
