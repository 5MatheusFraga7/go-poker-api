package models

import (
	"fmt"
	"sort"
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

func (p *Poker) SetCardPlayers(playerNumbers int) {

	p.Deck = NewDeck()

	displayPokerCards := DisplayPokerCards{PlayerHand: p.GetPlayerHand(), MachineHands: p.GetMachineHands(), TableCards: p.GetTableCards()}

	for i := 1; i < playerNumbers; i++ {
		p.Players = append(p.Players, Player{Id: strconv.Itoa(i), Name: fmt.Sprintf("Player %d", i+1), Hand: displayPokerCards.MachineHands[:2]})
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
		weight, greaterCard := p.CheckPlay(player.Hand, tableCards)
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
func (p *Poker) CheckPlay(playerHand []Card, tableCards []Card) (int, Card) {

	p.CheckPair(playerHand, tableCards)

	return 10, Card{}
}

func (p *Poker) CheckPair(playerHand []Card, tableCards []Card) (bool, int) {
	handValues := []int{}
	tableValues := []int{}

	for _, card := range playerHand {
		handValues = append(handValues, p.GetValueOfCards(card.Value))
	}

	for _, card := range tableCards {
		tableValues = append(tableValues, p.GetValueOfCards(card.Value))
	}

	// Confere se não tem um par na mão

	if handValues[0] == handValues[1] {
		return true, handValues[0]
	}

	// Confere se não tem um par na mão compondo com a mesa

	for _, handValue := range handValues {
		for _, tableValue := range tableValues {
			if handValue == tableValue {
				return true, handValue
			}
		}
	}

	return true, 0
}

func (p *Poker) CheckTwoPairs(playerHand []Card, tableCards []Card) bool {
	handValues := []int{}
	tableValues := []int{}
	pairsFound := 0

	for _, card := range playerHand {
		handValues = append(handValues, p.GetValueOfCards(card.Value))
	}

	for _, card := range tableCards {
		tableValues = append(tableValues, p.GetValueOfCards(card.Value))
	}

	// Confere se não tem um par na mão
	if handValues[0] == handValues[1] {
		pairsFound++
	}

	// Confere se não tem um par na mão compondo com a mesa

	for _, handValue := range handValues {
		for _, tableValue := range tableValues {
			if handValue == tableValue {
				pairsFound++
			}
		}
	}

	return pairsFound == 2
}

func (p *Poker) CheckThreeOfKind(playerHand []Card, tableCards []Card) (bool, int) {
	allValues := []int{}
	combinationFound := false

	foundedValue := 0

	for _, card := range playerHand {
		allValues = append(allValues, p.GetValueOfCards(card.Value))
	}

	for _, card := range tableCards {
		allValues = append(allValues, p.GetValueOfCards(card.Value))
	}

	sort.Ints(allValues)

	for i := 0; i <= len(allValues)-3; i++ {
		if allValues[i] == allValues[i+1] && allValues[i+1] == allValues[i+2] {
			combinationFound = true
			foundedValue = allValues[i]
		}
	}

	return combinationFound, foundedValue
}

func (p *Poker) CheckStraight(playerHand []Card, tableCards []Card) (bool, []int) {
	allValues := []int{}
	combinationFound := false

	for _, card := range playerHand {
		allValues = append(allValues, p.GetValueOfCards(card.Value))
	}

	for _, card := range tableCards {
		allValues = append(allValues, p.GetValueOfCards(card.Value))
	}

	sort.Ints(allValues)
	fmt.Println("Cartas em jogo: ", allValues)

	startValues := allValues[:5]
	midValues := allValues[1:6]
	lastValues := allValues[2:]

	if isArithmeticSequence(startValues) {
		return true, startValues
	}

	if isArithmeticSequence(midValues) {
		return true, midValues
	}

	if isArithmeticSequence(lastValues) {
		return true, lastValues
	}

	return combinationFound, nil
}

func (p *Poker) CheckFlush(playerHand []Card, tableCards []Card) bool {
	allSuits := []string{}
	startCombinationFound := false
	midCombinationFound := false
	lastCombinationFound := false

	combinationFound := false

	for _, card := range playerHand {
		allSuits = append(allSuits, card.Suit)
	}

	for _, card := range tableCards {
		allSuits = append(allSuits, card.Suit)
	}
	sort.Strings(allSuits)

	startValues := allSuits[:5]
	midValues := allSuits[1:6]
	lastValues := allSuits[2:]

	fmt.Println("Naipes em jogo: ", allSuits)

	for i := 0; i < len(startValues); i++ {
		if startValues[0] != startValues[i] {
			startCombinationFound = false
		}
	}

	for i := 0; i < len(midValues); i++ {
		if startValues[0] != midValues[i] {
			midCombinationFound = false
		}
	}

	for i := 0; i < len(lastValues); i++ {
		if startValues[0] != lastValues[i] {
			lastCombinationFound = false
		}
	}

	if startCombinationFound {
		combinationFound = true
	} else if midCombinationFound {
		combinationFound = true
	} else if lastCombinationFound {
		combinationFound = true
	}

	return combinationFound
}

func (p *Poker) CheckFullHouse(playerHand []Card, tableCards []Card) bool {

	hasPair, valuePairFound := p.CheckPair(playerHand, tableCards)
	hasThreeOfKind, valueThreeOfKindFound := p.CheckThreeOfKind(playerHand, tableCards)

	fmt.Println("Par encontrado em jogo: ", valuePairFound)
	fmt.Println("TRinca encontrado em jogo: ", valueThreeOfKindFound)

	bothValuesBiggerThanZero := valuePairFound > 0 && valueThreeOfKindFound > 0

	return hasPair && hasThreeOfKind && valuePairFound != valueThreeOfKindFound && bothValuesBiggerThanZero
}

func (p *Poker) CheckFourOfKind(playerHand []Card, tableCards []Card) bool {
	allValues := []int{}
	combinationFound := false

	// Adiciona os valores das cartas da mão do jogador ao slice allValues
	for _, card := range playerHand {
		allValues = append(allValues, p.GetValueOfCards(card.Value))
	}

	// Adiciona os valores das cartas da mesa ao slice allValues
	for _, card := range tableCards {
		allValues = append(allValues, p.GetValueOfCards(card.Value))
	}

	sort.Ints(allValues)

	for i := 0; i <= len(allValues)-4; i++ {
		// Verifica se os valores de 4 cartas consecutivas são iguais
		if allValues[i] == allValues[i+1] && allValues[i] == allValues[i+2] && allValues[i] == allValues[i+3] {
			combinationFound = true
			break // Sai do loop, já que encontramos uma quadra
		}
	}

	return combinationFound
}

func (p *Poker) CheckStraightFlush(playerHand []Card, tableCards []Card) bool {
	hasStraight, foundedValues := p.CheckStraight(playerHand, tableCards)

	if !hasStraight {
		return false
	}

	for _, value := range foundedValues {
		valueSuitCounter := 0

		for _, card := range append(playerHand, tableCards...) {
			if p.GetValueOfCards(card.Value) == value && card.Suit == playerHand[0].Suit {
				valueSuitCounter++
			}
		}

		if valueSuitCounter == 5 {
			return true
		}
	}

	return false
}

func (p *Poker) CheckRoyalFlush(playerHand []Card, tableCards []Card) bool {
	hasStraight, foundedValues := p.CheckStraight(playerHand, tableCards)

	if !hasStraight {
		return false
	}

	hasStraightFlush := p.CheckStraightFlush(playerHand, tableCards)

	if !hasStraightFlush {
		return false
	}

	for i := 10; i <= 14; i++ {
		if foundedValues[i] != i {
			return false
		}
	}

	return true
}

func isArithmeticSequence(sequence []int) bool {
	difference := sequence[1] - sequence[0]

	for i := 1; i < len(sequence)-1; i++ {
		if sequence[i+1]-sequence[i] != difference {
			return false
		}
	}

	return true
}

func (p *Poker) GetPokerCombinations() []Combination {
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

func (p *Poker) GetValueOfCards(value string) int {
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
