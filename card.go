package blackjack

type Card struct {
	Name string
	Values []int
}

var cardNameToValue = map[string][]int{
	"A":  []int{1, 11},
	"2":  []int{2},
	"3":  []int{3},
	"4":  []int{4},
	"5":  []int{5},
	"6":  []int{6},
	"7":  []int{7},
	"8":  []int{8},
	"9":  []int{9},
	"10": []int{10},
	"J":  []int{10},
	"Q":  []int{10},
	"K":  []int{10},
}

func NewCard(name string) Card {
	return Card{
		Name: name,
		Values: cardNameToValue[name],
	}
}