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

	poker := models.Poker{Deck: models.NewDeck()}
	poker.SetCardPlayers()

	fmt.Println("==============")
	fmt.Println("Players Hands:")
	fmt.Println("==============")
	for _, player := range poker.Players {
		fmt.Println("%s", player.Name)
		for _, card := range player.Hand {
			fmt.Println(card.Value+" de ", card.Suit)
		}
		fmt.Println("==============")
	}

	log.Printf(fmt.Sprintf("Server is running in %s:%s", config.Host, config.Port))
	log.Fatal(server.ListenAndServe())

}
