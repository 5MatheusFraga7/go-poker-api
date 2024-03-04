package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config := LoadConfig()
	router := mux.NewRouter()
	router.HandleFunc("/", HandlerFunc).Methods("GET")

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, config.Port),
		Handler: router,
	}

	// Inicializando o servidor HTTP
	log.Printf(fmt.Sprintf("Server is running in %s:%s", config.Host, config.Port))
	log.Fatal(server.ListenAndServe())
}

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"mensagem": "Olá, mundo!",
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
