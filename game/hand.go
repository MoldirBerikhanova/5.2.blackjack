package game

type hand struct {
	cards []card
}

func newHand() *hand {
	return &hand{}
}
func (h *hand) addCard(cardEx card) {
	if cardEx.value == 1 && h.getTotalValue()+11 <= 21 {
		cardEx = *newCard(cardEx.suit, 11)
	}
	h.cards = append(h.cards, cardEx)
}
func (h hand) getTotalValue() int {
	total := 0
	for i := 0; i < len(h.cards); i++ {
		total += h.cards[i].value
	}
	return total
}
