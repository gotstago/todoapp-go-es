package main

import (
	"math/rand"
	"time"
)

type Deck interface {
	Deal() Card
	Shuffle()
}

type GenericDeck struct {
	cards []Card
}

func (deck *GenericDeck) Shuffle() {
	rand.Seed(time.Now().Unix())
	for i := range deck.cards {
		j := rand.Intn(i + 1)
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	}
}

func (deck *GenericDeck) Deal() Card {
	if len(deck.cards) == 0 {
		return nil
	}
	card := deck.cards[0]
	deck.cards = deck.cards[1:]
	return card
}

type BlackJackDeck struct {
	*GenericDeck
}

func NewBlackJackDeck() *BlackJackDeck {
	bjdeck := &BlackJackDeck{&GenericDeck{}}

	for _, colour := range []string{"Kreuz", "Karo", "Herz", "Pik"} {
		for j := 1; j < 14; j++ {
			bjdeck.cards = append(bjdeck.cards, NewBlackJackCard(j, colour))
		}
	}

	return bjdeck
}
