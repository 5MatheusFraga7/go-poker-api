package models

import "fmt"

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

type Combination struct {
	Player      Player
	Weight      int
	GreaterCard Card
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

	combinations := []Combination{}

	for _, player := range players {
		weight, greaterCard := CheckCombination(player.Hand, tableCards)
		combinations = append(combinations, Combination{Player: player, Weight: weight, GreaterCard: greaterCard})
	}

	maxCombination := combinations[0]
	combinations = combinations[1:]

	for _, combination := range combinations {
		if combination.Weight > maxCombination.Weight {
			maxCombination = combination
		}
		if combination.Weight == maxCombination.Weight {
			maxCombination = CheckTieBreak(maxCombination, combination)
		}
	}

	// cards := p.CardsInGame
}

func CheckCombination(playerHand []Card, tableCards []Card) (int, Card) {
	return 10, Card{}
}

func CheckTieBreak(combinationA Combination, combinationB Combination) Combination {
	return combinationA
}
