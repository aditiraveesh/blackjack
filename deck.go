package blackjack

import "math/rand"

type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	return Deck{
		Cards: []Card{
			NewCard("A"), NewCard("2"), NewCard("3"),
			NewCard("4"), NewCard("5"), NewCard("6"), NewCard("7"),
			NewCard("8"), NewCard("9"), NewCard("10"), NewCard("J"),
			NewCard("Q"), NewCard("K"),
			NewCard("A"), NewCard("2"), NewCard("3"),
			NewCard("4"), NewCard("5"), NewCard("6"), NewCard("7"),
			NewCard("8"), NewCard("9"), NewCard("10"), NewCard("J"),
			NewCard("Q"), NewCard("K"),
			NewCard("A"), NewCard("2"), NewCard("3"),
			NewCard("4"), NewCard("5"), NewCard("6"), NewCard("7"),
			NewCard("8"), NewCard("9"), NewCard("10"), NewCard("J"),
			NewCard("Q"), NewCard("K"),
			NewCard("A"), NewCard("2"), NewCard("3"),
			NewCard("4"), NewCard("5"), NewCard("6"), NewCard("7"),
			NewCard("8"), NewCard("9"), NewCard("10"), NewCard("J"),
			NewCard("Q"), NewCard("K")}}
}

func (deck *Deck) Shuffle() {
	for index := range deck.Cards {
    	randomIndex := rand.Intn(index + 1)
    	deck.Cards[index], deck.Cards[randomIndex] = deck.Cards[randomIndex], deck.Cards[index]
	}
}

//TODO handle empty deck
func (deck *Deck) Deal() Card {
	firstCard := deck.Cards[0]
	deck.Cards = deck.Cards[1:]
	return firstCard
}

