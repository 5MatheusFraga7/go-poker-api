package routes

import (
	"example/internal/api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() http.Handler {
	router := mux.NewRouter()
	router.Use(LoggerMiddleware)
	router.NotFoundHandler = http.HandlerFunc(HandleNotFound)
	p := handlers.PokerHandler{}

	// Routes
	router.HandleFunc("/get_new_round", p.NewRound).Methods("GET")

	return router
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Recebida solicitação para %s %s com parâmetros %v", r.Method, r.URL.Path, r.URL.Query())

		next.ServeHTTP(w, r)
	})
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("Solicitação não encontrada para %s %s", r.Method, r.URL.Path)
	http.Error(w, "Endpoint não encontrado", http.StatusNotFound)
}
