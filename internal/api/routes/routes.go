package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeFunc).Methods("GET")

	return router
}

func HomeFunc(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"mensagem": "Olá, mundooooooooooooooooo!",
	}

	// Convertendo o mapa para JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Erro ao gerar JSON", http.StatusInternalServerError)
		return
	}

	// Definindo o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")

	// Escrevendo o JSON na resposta
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Erro ao escrever resposta", http.StatusInternalServerError)
		return
	}
}
