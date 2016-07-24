package blackjack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDeckShuffleShouldShuffleAllCards(t *testing.T) {
	deck := NewDeck()

	assert.Equal(t, len(deck.Cards), 52, "No of cards in shuffled deck should be 52")

	deck.Shuffle()

	assert.Equal(t, len(deck.Cards), 52, "No of cards in shuffled deck should be 52")
}

func TestDealCardShouldReturnACardAndRemoveItFromDeck(t * testing.T) {
	deck := NewDeck()

	assert.Equal(t, deck.Deal().Name, "A", "Card should be A")
	assert.Equal(t, deck.Deal().Name, "2", "Card should be 2")
	assert.Equal(t, deck.Deal().Name, "3", "Card should be 3")
	assert.Equal(t, deck.Deal().Name, "4", "Card should be 4")
	assert.Equal(t, deck.Deal().Name, "5", "Card should be 5")
}