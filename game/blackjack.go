package game

import "fmt"

type blackjack struct {
	// Это структура игры, она в качестве филдов должна содержать
	// колоду, игроков и дилера
	deckEx   deck
	dealerEx dealer
	playerEx player
}

func NewBlackjack() *blackjack {
	deckVar := newDeck()
	dealerVar := newDealer()
	playerVar := newPlayer()

	return &blackjack{
		deckEx:   deckVar,
		dealerEx: dealerVar,
		playerEx: playerVar,
	}
}
func (b *blackjack) Play() {
	/*
		Для начала игры нам нужна колода карт, игрок и дилер.

		Действия в игре выполняются в определенной последовательности.
		В начале раунда игроку выдается две карты, полученные карты нужно отобразить в консоль.
		Дилеру так же выдается две карты, но первая карта как бы кладется рубашкой вверх, поэтому ее номинал не нужно выводить в консоль.

		Игрок принимает решение взять ему еще одну карту или нет. Добирать карту можно неограниченное количество раз. Цель - победить дилера, не превысив 21 очка.

		В конце игры вывести в консоль сумму руки игрока и дилера и определить победителя.
	*/
	b.deckEx.shuffleCards()
	for i := 0; i < 2; i++ {
		card := b.deckEx.dealCard()
		b.playerEx.addCard(card)
	}
	for i := 0; i < 2; i++ {
		card := b.deckEx.dealCard()
		b.dealerEx.addCard(card)
	}

	for {
		fmt.Printf("Взять еще одну карту (д/н)?")
		var input string
		fmt.Scan(&input)

		if input == "н" {
			break
		}
		card := b.deckEx.dealCard()
		b.playerEx.addCard(card)

		//почситать количество очков руки
		if b.playerEx.isBusted() {
			b.whoIsWinner()
			return
		}
	}
	for {
		total := b.dealerEx.getHandTolalValue()
		if total <= 17 {
			b.dealerEx.addCard(b.deckEx.dealCard())
		} else {

			b.whoIsWinner()
			return
		}
	}
}

func (b blackjack) whoIsWinner() {
	fmt.Printf(
		"\n Сумма руки игрока - %d очков. Сумма руки дилера - %d очков\n",
		b.playerEx.getHandTolalValue(),
		b.dealerEx.getHandTolalValue())

	if b.playerEx.isBusted() {
		fmt.Printf("Игрок сгорел! Дилер выигрывает\n")
		return
	}
	if b.dealerEx.isBusted() {
		fmt.Printf("Дилер сгорел! Дилер выигрывает")
		return
	}
	if b.playerEx.getHandTolalValue() > b.dealerEx.getHandTolalValue() {
		fmt.Printf("Игрок выиграл")
	} else {
		fmt.Printf("Игрок проиграл")
	}
}
