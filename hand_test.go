package blackjack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHandScoreWithNumberValues(t *testing.T) {
	hand := Hand{Cards: []string{"2", "3", "5", "7"}}

	assert.Equal(t, hand.Scores(), []int{17}, "Score should equal the total of card values in hand")
}

func TestHandScoreWithJack(t *testing.T) {
	hand := Hand{Cards: []string{"J", "3"}}

	assert.Equal(t, hand.Scores(), []int{13}, "Score should equal the total of card values in hand")
}

func TestHandScoreWithQueen(t *testing.T) {
	hand := Hand{Cards: []string{"Q", "5"}}

	assert.Equal(t, hand.Scores(), []int{15}, "Score should equal the total of card values in hand")
}

func TestHandScoreWithKing(t *testing.T) {
	hand := Hand{Cards: []string{"K", "6"}}

	assert.Equal(t, hand.Scores(), []int{16}, "Score should equal the total of card values in hand")
}

func TestHandScoreWithAce(t *testing.T) {
	hand := Hand{Cards: []string{"A", "3"}}
	assert.Equal(t, hand.Scores(), []int{4, 14}, "Score should equal the total of card values in hand")

	hand = Hand{Cards: []string{"A", "A"}}
	assert.Equal(t, hand.Scores(), []int{2, 12, 22}, "Score should equal the total of card values in hand")

	hand = Hand{Cards: []string{"A", "A", "A", "2"}}
	assert.Equal(t, hand.Scores(), []int{5, 15, 25, 35}, "Score should equal the total of card values in hand")
}
