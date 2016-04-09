package game

//Game model
type Game struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
    Stacks    []Stack `json:"stacks"`
}

//Card will capture the details of a playing card
type Card struct {
    Rank string `json:"rank"`
    Suit string `json:"suit"`
}

//Stack will hold a collection of cards and a name
type Stack struct{
    Name string `json:"name"`
    Cards []Card `json:"cards"`
}