package game

import (
	"math/rand"
	"time"
)

type deck struct {
	cards []card
}

func newDeck() deck {
	d := deck{}
	d.cards = d.generateCards()

	return d
}

func (d deck) generateCards() []card {
	var cardSuits []string = []string{"Червы", "Пики", "Бубны", "Крести"}
	var cardValues []string = []string{"Туз", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Джек", "Королева", "Король"}

	var cards []card
	for _, s := range cardSuits {
		for _, v := range cardValues {
			//	card := fmt.Sprintf("%s %s", s, v)
			cardValue := getCardValue(v)
			cardEx := newCard(s, cardValue)
			cards = append(cards, *cardEx)
		}
	}

	// Реализовать метод генерации колоды.
	// Для каждой масти нужно создать карту с необходимым номиналом.
	//
	// В данном случае картой считаем строку в формате "{масть} {номинал}", к примеру, "Червы 2"

	// ВАШ КОД ТУТ...

	return cards
}

func getCardValue(cardValue string) int {
	if cardValue == "Туз" {
		return 1
	} else if cardValue == "2" {
		return 2
	} else if cardValue == "3" {
		return 3
	} else if cardValue == "4" {
		return 4
	} else if cardValue == "5" {
		return 5
	} else if cardValue == "6" {
		return 6
	} else if cardValue == "7" {
		return 7
	} else if cardValue == "8" {
		return 8
	} else if cardValue == "9" {
		return 9
	} else if cardValue == "10" {
		return 10
	} else if cardValue == "Джек" {
		return 10
	} else if cardValue == "Королева" {
		return 10
	} else if cardValue == "Король" {
		return 10
	} else {
		return -1
	}
}

func (d deck) shuffleCards() {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)

	// Для того, чтобы перемешать колоду, мы пробегаемся по всем картам используя `for` цикл
	// и меняем карты местами.
	// Для определения нового места (индекса) для карты, используется генерация случайного числа.

	for i := 0; i < len(d.cards); i++ {
		newIndex := rand.Intn(len(d.cards))
		temp := d.cards[i]
		d.cards[i] = d.cards[newIndex]
		d.cards[newIndex] = temp

		/*
			Используя `newIndex` реализуйте логику перестановки двух карт местами.
			В каждой итерации цикла у нас есть
			- i - текущий индекс (позиция) карты
			- d.cards[i] - текущая карта
			- newIndex - новый индекс карты, куда она должны быть переставлена
			- d.cards[newIndex] - карта с которой нужно поменяться местами

			Для того, чтобы переставить два значения местами в массиве, мы можем использовать
			дополнительную переменную, в которую сохраним временное значение.

			Пример. Есть два числа, а = 10, b = 5, мы хотим переставить их местами, то есть чтобы a = 5, b = 10.
			a := 10
			b := 5

			c := a
			a = b
			b = c
		*/

		// ВАШ КОД ТУТ...
	}
}

func (d *deck) dealCard() card {
	// Взять карту с колоды - значит взять самую первую карту из нее, то есть взять карту по индексом 0.
	card := d.cards[0]

	// Так как мы взяли вернхнюю карту, мы должны удалить ее из колоды.
	// Используя оператор `slice` [1:] мы из нашего начального слайса берем все элементы кроме самого первого.
	d.cards = d.cards[1:]

	return card
}
