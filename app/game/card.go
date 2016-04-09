package main

import "fmt"

type Card interface {
	getValue() int
	getColour() string
	isAce() bool
}

type BlackJackCard struct {
	Value  int
	Colour string
}

func (bjcard *BlackJackCard) getValue() int {
	if bjcard.Value >= 11 && bjcard.Value <= 13 {
		return 10
	}
	return bjcard.Value
}

func (bjcard *BlackJackCard) getColour() string {
	return bjcard.Colour
}

func (bjcard *BlackJackCard) String() string {
	return fmt.Sprintf("Wert %d - Farbe %s (realValue: %d)", bjcard.getValue(), bjcard.getColour(), bjcard.Value)
}

func (bjcard *BlackJackCard) isAce() bool {
	return bjcard.Value == 1
}

func NewBlackJackCard(value int, colour string) Card {
	return &BlackJackCard{Value: value, Colour: colour}
}
