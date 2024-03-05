package main

import (
	"example/internal/api/routes"
	"example/internal/models"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config := LoadConfig()
	router := routes.SetupRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, config.Port),
		Handler: router,
	}

	deck := models.NewDeck()
	deck.Shuffle()

	fmt.Printf("deck: %v\n", deck)

	log.Printf(fmt.Sprintf("Server is running in %s:%s", config.Host, config.Port))
	log.Fatal(server.ListenAndServe())

}
