package models

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Deck []Card

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
}

func NewDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Espadas", "Ouros", "Paus", "Copas"}
	keyCards := []string{"Ás", "Valete", "Dama", "Rei"}

	for _, suit := range cardSuits {
		for i := 2; i <= 10; i++ {
			cards = append(cards, Card{Value: strconv.Itoa(i), Suit: suit})

		}
		for _, keyCard := range keyCards {
			cards = append(cards, Card{Value: string(keyCard), Suit: suit})
		}
	}

	return cards

}

func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) Shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
