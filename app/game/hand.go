package main

import (
	"fmt"
	"math"
)

type Hand interface {
	AddCard(Card)
	Points() int
	PrintHand()
}

type BlackJackHand struct {
	cards []Card
}

func (h *BlackJackHand) AddCard(c Card) {
	h.cards = append(h.cards, c)
}

func (h *BlackJackHand) Points() int {
	scores := h.possibleScores()
	if len(scores) == 0 {
		return 0
	}
	maxUnder := math.MinInt64
	minOver := math.MaxInt64
	for _, score := range scores {
		if score > 21 && score < minOver {
			minOver = score
		} else if score <= 21 && score > maxUnder {
			maxUnder = score
		}
	}
	if maxUnder == math.MinInt64 {
		return minOver
	}
	return maxUnder
}

func (h *BlackJackHand) possibleScores() []int {
	scores := make([]int, 0)
	if len(h.cards) == 0 {
		return scores
	}
	for _, c := range h.cards {
		scores = h.addToScoreList(c, scores)
	}
	return scores
}

func (h *BlackJackHand) addToScoreList(c Card, scores []int) []int {
	if len(scores) == 0 {
		scores = append(scores, 0)
	}
	length := len(scores)
	for i := 0; i < length; i++ {
		currentScore := scores[i]
		scores[i] = currentScore + c.getValue()
		if c.isAce() {
			scores = append(scores, currentScore+11)
		}
	}
	return scores
}

func (h *BlackJackHand) PrintHand() {
	for _, card := range h.cards {
		fmt.Println(card)
	}
}

func NewBlackJackHand() Hand {
	return &BlackJackHand{}
}
