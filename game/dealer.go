package game

import "fmt"

type dealer struct {
	handEx *hand
}

func newDealer() dealer {
	return dealer{newHand()}
}

func (d *dealer) addCard(cardEx card) {
	// Реализовать логику добавления карты игроку
	d.handEx.addCard(cardEx)
	if len(d.handEx.cards) == 1 {
		fmt.Println("дилер взял карту")
	} else {
		fmt.Printf("диллер взял карту %s %d\n", cardEx.suit, cardEx.value)
	}
}

func (d *dealer) getHandTolalValue() int {
	return d.handEx.getTotalValue()
}
