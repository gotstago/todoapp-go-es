package main

import "fmt"

type Game interface {
	Start()
}

type BlackJackGame struct {
	player *Player
	dealer *Dealer
	deck   Deck
}

func (game *BlackJackGame) DealInitial() {
	game.player.AddCard(game.deck.Deal())
	game.dealer.AddCard(game.deck.Deal())
	game.player.AddCard(game.deck.Deal())
	game.dealer.AddCard(game.deck.Deal())
}

func (game *BlackJackGame) Start() {
	game.DealInitial()
	var input string
	fmt.Println("Type \"q\" if you dont want more cards, Enter otherwise")
	fmt.Printf("Your Score %d \t Dealer Score %d\n", game.player.Points(), game.dealer.Points())
	for !game.player.Busted() && !game.dealer.Busted() {
		fmt.Scanf("%s", &input)
		if input == "q" {
			break
		}
		game.player.Play(game.deck)
		game.dealer.Play(game.deck)
		fmt.Printf("Your Score %d \t Dealer Score %d\n", game.player.Points(), game.dealer.Points())
	}

	// play dealer
	// he must play until he hits more then 16 points
	dealerpoints := game.dealer.Points()
	for !game.dealer.Busted() {
		game.dealer.Play(game.deck)
		if dealerpoints == game.dealer.Points() {
			break
		} else {
			dealerpoints = game.dealer.Points()
		}
		fmt.Printf("Your Score %d \t Dealer Score %d\n", game.player.Points(), game.dealer.Points())
	}

	if game.player.Busted() && game.dealer.Busted() {
		fmt.Println("Its a draw - both busted")
	} else if game.player.Busted() {
		fmt.Println("You lose")
	} else if game.dealer.Busted() {
		fmt.Println("You win!")
	} else {
		if game.player.Points() > game.dealer.Points() {
			fmt.Println("You win!")
		} else {
			fmt.Println("You lose")
		}
	}
	fmt.Println("\nYour Hand")
	game.player.PrintHand()
	fmt.Println("Dealers Hand")
	game.dealer.PrintHand()
}

func NewBlackJackGame() *BlackJackGame {
	deck := NewBlackJackDeck()
	deck.Shuffle()
	return &BlackJackGame{player: NewPlayer(), dealer: NewDealer(), deck: deck}
}
