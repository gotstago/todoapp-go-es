package main

type Player struct {
	Hand
}

func (p *Player) Busted() bool {
	points := p.Points()
	if points > 21 {
		return true
	}
	return false
}

func (p *Player) Play(deck Deck) {
	p.AddCard(deck.Deal())
}

func NewPlayer() *Player {
	return &Player{NewBlackJackHand()}
}

type Dealer struct {
	*Player
	hitUntil int
}

func (d *Dealer) Play(deck Deck) {
	if d.Points() < d.hitUntil {
		d.Player.Play(deck)
	}
}

func NewDealer() *Dealer {
	return &Dealer{Player: NewPlayer(), hitUntil: 16}
}
