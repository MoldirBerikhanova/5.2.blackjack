package game

type card struct {
	suit  string
	value int
}

func newCard(suit string, value int) *card {
	return &card{suit, value}
}

// func (h *hand) addCard(cardEx card){
// 	h.cards = append(h.cards, cardEx)
// }