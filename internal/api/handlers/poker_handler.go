package handlers

import (
	"encoding/json"
	"example/internal/models"
	"log"
	"net/http"
	"strconv"
)

type PokerHandler struct {
}

func (p *PokerHandler) HandleInternalServerError(w http.ResponseWriter, r *http.Request) {
	log.Printf("Erro interno do servidor ao processar solicitação para %s %s com parâmetros %v", r.Method, r.URL.Path, r.URL.Query())
	http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
}

func (p *PokerHandler) NewRound(w http.ResponseWriter, r *http.Request) {

	playerNumbers, err := strconv.Atoi(r.URL.Query().Get("playerNumbers"))

	poker := models.Poker{}
	poker.SetCardPlayers(playerNumbers)

	data := map[string]interface{}{
		"players":    poker.Players,
		"tableCards": poker.TableCards,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		p.HandleInternalServerError(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonData)
	if err != nil {
		p.HandleInternalServerError(w, r)
		return
	}
}
func (p *PokerHandler) GetWinner(w http.ResponseWriter, r *http.Request) {
	var data models.PokerData
	var poker models.Poker

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	poker.AvailablePlayers = data.Players
	poker.TableCards = data.TableCards

	winner := poker.GetWinner()

	dataResponse := map[string]interface{}{
		"winner":      winner.Player,
		"combination": winner.Combination,
	}
	jsonData, err := json.Marshal(dataResponse)

	// Resposta de exemplo
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
