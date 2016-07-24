package blackjack

type Hand struct {
	Cards []Card
}

func (hand Hand) Scores() []int {
	possibleScores := []int{0}

	for _, card := range hand.Cards {
		possibleScores = hand.duplicateScoresWithCardValuesAdded(possibleScores, card)
	}

	return possibleScores
}

func (hand Hand) duplicateScoresWithCardValuesAdded(possibleScores []int, card Card) []int {
	scores := []int{}

	for _, cardValue := range card.Values {
		for _, possibleScore := range possibleScores {
			if(!hand.contains(scores, possibleScore + cardValue)) {
				scores = append(scores, possibleScore + cardValue)	
			}
		}
	}

	return scores
}

func (hand Hand) contains(array []int, item int) bool {
	for _, value := range array {
		if value == item {
			return true
		}
	}
	return false
} 












