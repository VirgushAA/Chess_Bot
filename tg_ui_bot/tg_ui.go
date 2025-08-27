package main

import (
	. "Chess_Bot/core"
	"encoding/json"
	"github.com/google/uuid"
	// "log"
	"net/http"
	"sync"
)

var (
	games = make(map[string]*Game)
	mutex sync.Mutex
)

type MoveRequest struct {
	GameID string `json:"gameId"`
	Move   string `json:"move"`
}

func newGameHandler(w http.ResponseWriter, r *http.Request) {
	g := &Game{}
	g.GameStart()

	id := uuid.New().String()

	mutex.Lock()
	games[id] = g
	mutex.Unlock()

	json.NewEncoder(w).Encode(map[string]any{
		"gameId": id,
		"state":  g.GameState,
	})
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	var request MoveRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mutex.Lock()
	g, ok := games[request.GameID]
	mutex.Unlock()

	if !ok {
		http.Error(w, "Game not found", http.StatusNotFound)
	}

	g.PlayATurn(request.Move)

	json.NewEncoder(w).Encode(map[string]any{
		"state": g.GameState,
	})

}

func main() {
	http.HandleFunc("/newgame", newGameHandler)
	http.HandleFunc("/move", moveHandler)
	http.ListenAndServe(":8080", nil)
}
