package tguibot

import (
	. "Chess_Bot/core"
	"encoding/json"
	// "log"
	"net/http"
	"sync"
)

var (
	games = make(map[string]*Game)
	mutex sync.Mutex
)

type MoveRequest struct {
	GameID string
	Move   string
}

func newGameHandler(w http.ResponseWriter, r *http.Request) {

}

func moveHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/newgame", newGameHandler)
	http.HandleFunc("/move", moveHandler)
	http.ListenAndServe(":8080", nil)
}
