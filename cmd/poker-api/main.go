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
	fmt.Println("Players Hand:")
	fmt.Println("==============")

	playerHand := poker.GetPlayerHand()

	for _, card := range playerHand {
		fmt.Println(card.Value+" de ", card.Suit)
	}

	fmt.Println("==============")
	fmt.Println("Table Cards:")
	fmt.Println("==============")

	tableCards := poker.GetTableCards()

	for _, card := range tableCards {
		fmt.Println(card.Value+" de ", card.Suit)
	}
	fmt.Println("==============")

	if poker.CheckPair(playerHand, tableCards) {
		fmt.Println("tEMOS UM PAR!")
	} else {
		fmt.Println("sem par!")
	}

	if poker.CheckTwoPairs(playerHand, tableCards) {
		fmt.Println("tEMOS DOIS PARES!")
	}

	log.Printf(fmt.Sprintf("Server is running in %s:%s", config.Host, config.Port))
	log.Fatal(server.ListenAndServe())

}
