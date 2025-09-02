package main

import (
	. "Chess_Bot/ai"
	. "Chess_Bot/core"
	"encoding/json"
	"github.com/google/uuid"
	// "log"
	"fmt"
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
	for id := range games {
		fmt.Println("GameID:", id)
		fmt.Println()
	}
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

	print(request.Move)
	g.PlayATurn(request.Move)

	json.NewEncoder(w).Encode(map[string]any{
		"state": g.GameState,
	})
}

func moveHandlerAI(w http.ResponseWriter, r *http.Request) {
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

	MakeAMoveAI(g)

	json.NewEncoder(w).Encode(map[string]any{
		"state": g.GameState,
	})
}

func endGameHandler(w http.ResponseWriter, r *http.Request) {
	var request MoveRequest
	if err := json.NewDecoder(r.Body).Decode((&request)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mutex.Lock()
	_, ok := games[request.GameID]
	if ok {
		delete(games, request.GameID)
	}
	mutex.Unlock()

	if !ok {
		http.Error(w, "Game not found", http.StatusNotFound)
	}
	// for id, g := range games {
	// 	fmt.Printf("GameID: %s, Game state: %+v\n", id, g)
	// }
	for id := range games {
		fmt.Println("GameID:", id)
		fmt.Println()
	}
	if len(games) == 0 {
		fmt.Print("No active games.\n")
	}
}

func main() {
	http.HandleFunc("/newgame", newGameHandler)
	http.HandleFunc("/move", moveHandler)
	http.HandleFunc("/endgame", endGameHandler)
	http.HandleFunc("/moveAI", moveHandlerAI)
	http.ListenAndServe(":8080", nil)
}
