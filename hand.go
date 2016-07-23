package blackjack

import "strconv"

type Hand struct {
	Cards []string
}

func (hand Hand) Scores() []int {
	possibleScores := []int{}
	noOfAces := 0
	score := 0

	for _, card := range hand.Cards {
		if card == "A" {
			noOfAces ++
		}

		score += hand.getScoreOfNonAce(card)
	}

	if noOfAces == 0 {
		return []int{score}
	}

	for i := 0; i <= noOfAces; i ++ {
		possibleScores = append(possibleScores, score + ((i  * 10) + noOfAces))
	}

 	return possibleScores
}

func (hand Hand) getScoreOfNonAce(card string) int {
	if card == "J" || card == "Q" || card == "K" {
		return 10
	} else if card != "A" {
		value, _ := strconv.Atoi(card)
		return value
	} else {
		return 0
	}
}