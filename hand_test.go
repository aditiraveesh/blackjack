package blackjack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHandScoreWithNumberValues(t *testing.T) {
	two := NewCard("2")
	three := NewCard("3")
	five := NewCard("5")
	seven := NewCard("7")

	hand := Hand{Cards: []Card{two, three, five, seven}}

	assert.Equal(t, hand.Scores(), []int{17}, "Score should equal the total of card values in hand")
}

func TestHandScoreWithJack(t *testing.T) {
	jack := NewCard("J")
	four := NewCard("4")

	hand := Hand{Cards: []Card{jack, four}}

	assert.Equal(t, hand.Scores(), []int{14}, "Score should equal the total of card values in hand")
}

func TestHandScoreWithQueen(t *testing.T) {
	queen := NewCard("Q")
	five := NewCard("5")

	hand := Hand{Cards: []Card{queen, five}}

	assert.Equal(t, hand.Scores(), []int{15}, "Score should equal the total of card values in hand")
}

func TestHandScoreWithKing(t *testing.T) {
	king := NewCard("K")
	six := NewCard("6")

	hand := Hand{Cards: []Card{king, six}}

	assert.Equal(t, hand.Scores(), []int{16}, "Score should equal the total of card values in hand")
}

func TestHandScoreWithAce(t *testing.T) {
	ace := NewCard("A")
	three := NewCard("3")

	hand := Hand{Cards: []Card{ace, three}}

	assert.Equal(t, hand.Scores(), []int{4, 14}, "Score should equal the total of card values in hand")

	hand = Hand{Cards: []Card{ace, ace}}
	assert.Equal(t, hand.Scores(), []int{2, 12, 22}, "Score should equal the total of card values in hand")

	hand = Hand{Cards: []Card{ace, ace, ace, three}}
	assert.Equal(t, hand.Scores(), []int{6, 16, 26, 36}, "Score should equal the total of card values in hand")
}
