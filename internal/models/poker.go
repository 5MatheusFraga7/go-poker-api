package models

import (
	"fmt"
	"strconv"
)

type Poker struct {
	Deck             Deck
	Players          []Player
	AvailablePlayers []Player
	TableCards       []Card
}

type Player struct {
	Name  string
	Id    string
	Hand  []Card
	Money float64
}

type DisplayPokerCards struct {
	PlayerHand   []Card
	MachineHands []Card
	TableCards   []Card
}

type Play struct {
	Player      Player
	Weight      int
	GreaterCard Card
}

type Combination struct {
	Name   string
	Weight int
}

func (p *Poker) SetCardPlayers() {

	displayPokerCards := DisplayPokerCards{PlayerHand: p.GetPlayerHand(), MachineHands: p.GetMachineHands(), TableCards: p.GetTableCards()}

	for i := 1; i <= 4; i++ {
		p.Players = append(p.Players, Player{Name: fmt.Sprintf("Player %d", i+1), Hand: displayPokerCards.MachineHands[:2]})
		displayPokerCards.MachineHands = displayPokerCards.MachineHands[2:]
	}

	p.Players = append(p.Players, Player{Name: fmt.Sprintf("Player 1"), Hand: displayPokerCards.PlayerHand})
}

func (p *Poker) GetPlayerHand() []Card {
	playerHand := p.Deck[:2]

	p.Deck = p.Deck[2:]

	return playerHand
}

func (p *Poker) GetMachineHands() []Card {
	machineHands := p.Deck[:8]
	p.Deck = p.Deck[8:]

	return machineHands
}

func (p *Poker) GetTableCards() []Card {

	tableCards := p.Deck[:5]
	p.Deck = p.Deck[5:]

	p.TableCards = tableCards
	return tableCards
}

func (p *Poker) RemoveAvailablePlayer(playerIndex int) {
	if playerIndex < 0 || playerIndex >= len(p.AvailablePlayers) {
		fmt.Println("Índice inválido")
		return
	}

	p.AvailablePlayers = append(p.AvailablePlayers[:playerIndex], p.AvailablePlayers[playerIndex+1:]...)
}

func (p *Poker) GetWinner() {

	players := p.AvailablePlayers
	tableCards := p.TableCards

	plays := []Play{}

	for _, player := range players {
		weight, greaterCard := CheckPlay(player.Hand, tableCards)
		plays = append(plays, Play{Player: player, Weight: weight, GreaterCard: greaterCard})
	}

	maxPlay := plays[0]
	plays = plays[1:]

	for _, play := range plays {
		if play.Weight > maxPlay.Weight {
			maxPlay = play
		}
		if play.Weight == maxPlay.Weight {
			maxPlay = CheckTieBreak(maxPlay, play)
		}
	}

	// cards := p.CardsInGame
}

func GetPokerCombinations() []Combination {
	combinations := []Combination{
		{Name: "One Pair", Weight: 1},
		{Name: "Two Pairs", Weight: 2},
		{Name: "Three of a Kind", Weight: 3},
		{Name: "Straight", Weight: 4},
		{Name: "Flush", Weight: 5},
		{Name: "Full House", Weight: 6},
		{Name: "Quadra", Weight: 7},
		{Name: "Straight Flush", Weight: 8},
		{Name: "Royal Flush", Weight: 9},
	}
	return combinations
}

func CheckPlay(playerHand []Card, tableCards []Card) (int, Card) {

	CheckPair(playerHand, tableCards)

	return 10, Card{}
}

func CheckPair(playerHand []Card, tableCards []Card) bool {
	// combinations := GetPokerCombinations()
	for _, card := range playerHand {
		GetValueOfKeyCard(card.Value)
	}

	return false
}

func GetValueOfKeyCard(value string) int {
	response, _ := strconv.Atoi(value)
	switch value {
	case "Ás":
		return 14
	case "Valete":
		return 11
	case "Dama":
		return 12
	case "Rei":
		return 13
	default:
		return response
	}
}

// Retorna o desempate

func CheckTieBreak(PlayA Play, PlayB Play) Play {
	return PlayA
}
