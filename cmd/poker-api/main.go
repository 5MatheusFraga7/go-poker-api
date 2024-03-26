package main

import (
	"example/internal/api/routes"
	"example/internal/models"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	config := LoadConfig()
	router := routes.SetupRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, config.Port),
		Handler: router,
	}

	for {
		runProgram()
		time.Sleep(1 * time.Second) // Espera 2 segundos antes de executar novamente
	}

	log.Printf(fmt.Sprintf("Server is running in %s:%s", config.Host, config.Port))
	log.Fatal(server.ListenAndServe())
}

func runProgram() {
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

	checkPair, _ := poker.CheckPair(playerHand, tableCards)

	if checkPair {
		fmt.Println("tEMOS UM PAR!")
	} else {
		fmt.Println("sem par!")
	}
	checkThreeOfKind, _ := poker.CheckThreeOfKind(playerHand, tableCards)

	if checkThreeOfKind {
		fmt.Println("tEMOS uma TRINCA!!!")
	}

	if poker.CheckStraight(playerHand, tableCards) {
		fmt.Println("tEMOS uma SEQUENCIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA!!!")
	}

	if poker.CheckFlush(playerHand, tableCards) {
		fmt.Println("tEMOS um FLUUUUSHHH!!!!!!!!")
	}

	if poker.CheckFullHouse(playerHand, tableCards) {
		fmt.Println("tEMOS um FULL HOUSE !!!!!!!!")
	}

	if poker.CheckFourOfKind(playerHand, tableCards) {
		fmt.Println("tEMOS um FOUR OF KIND !!!!!!!!")
	}

}
