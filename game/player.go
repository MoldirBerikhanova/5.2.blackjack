package game

import "fmt"

type player struct {
	handEx *hand
}

func newPlayer() player {
	return player{newHand()}
}

func (g *player) addCard(cardEx card) {
	g.handEx.addCard(cardEx)
	fmt.Printf("игрок Моля взял карту %s %d\n", cardEx.suit, cardEx.value)
}

func (g *player) getHandTolalValue() int {
	return g.handEx.getTotalValue()
}

func (g *player) isBusted() bool{
	return g.handEx.getTotalValue() > 21
}

func (g *dealer) isBusted() bool{
	return g.handEx.getTotalValue() > 21
}