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
	var pokerData models.PokerData

	// Decodifica o JSON do corpo da solicitação para a struct PokerData

	err := json.NewDecoder(r.Body).Decode(&pokerData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ADICIONAR AQUI A CHAMDA PARA O MÉTDODO GET WINNER QUE AINDA NÃO FOI IMPLEMENTADO EM POKER.GO

	// players := pokerData.Players
	// tableCards := pokerData.TableCards

	//gerando resposta

	data := map[string]interface{}{
		"players":    pokerData.Players,
		"tableCards": pokerData.TableCards,
	}
	jsonData, err := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
}
