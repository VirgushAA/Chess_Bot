package main

import (
	. "Chess_Bot/core"
	// "fmt"
)

var ()

func main() {
	board := NewBoard()
	board.Print()

move := GetUserInput()
if move in game.LegalMoves() {
    game.MakeMove(move)
    game.Print()
    engineMove := SearchBestMove(game)
    game.MakeMove(engineMove)
}

}