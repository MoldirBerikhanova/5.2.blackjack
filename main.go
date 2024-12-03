package main

import (
	"blackjack/game"
	
	//"strings"
	//"time"
)



func main() {

	b := game.NewBlackjack()
	b.Play()
}


	///////////////////////// Cоздание Колоды/////////////////////

	

// p := Player{name: "Моля", values:v}
	//fmt.Println(p)
	/////создание колоды
	//p := NewPlayer{name:"Моля", values:PlayersValue[i]}/// это для функции в итеррации

	//var Players Player

	//var Play Players
	//fmt.Println("Русская колода:", deck)
	//fmt.Println(dealCards(deck))

	// source := rand.NewSource(time.Now().UnixNano())
	// rand := rand.New(source)
	// rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	//PlayersValue := dealCards(deck)

	// Output the dealt cards
	//fmt.Println(PlayersValue)
	//fmt.Println(dealCards(deck))

	
// var play Player
// for _, v := range PlayersValue {
// 	play = append(play, Player{values: v})
// }

// var result []string
// for _, v := range numb {
// 	if x, found := R[v]; found {
// 		result = append(result, x)
// 	} else {
// 		result = append(result, "")
// 	}
// }

// ////////////надо завернутьв функцию/////////////

// 	}
// 	}

// var DilersValue []string
// for i := 0; i <= 5; i++ {
// 	index := rand.Intn(len(deck))
// 	value := deck[index]
// 	DilersValue = append(DilersValue, fmt.Sprintf("%s", value))
// 	deck = append(deck[:index], deck[index+1:]...) ///////////%s переводит
// 	fmt.Printf(" "+"#%v"+" "+"Карта"+" "+"%s\n", i, DilersValue[i]+" ")

// }

//Residue := len(deck)-len(PlayersValue)-len(DilersValue)
//fmt.Println(deck)

// for i := 0; i < len(PlayersValue); i++ {
// 	if PlayersValue[i] != DilersValue[i] {
// 		fmt.Printf("Карта %s: Player's value = %s, Dealer's value = %s\n", i, PlayersValue[i], DilersValue[i])
// 	}
// }
//fmt.Println("Дилеры", DilersValue)

////остаток колоды/////

////var PlayerValue []string
//var Residue []string

////var PlayerValue []string
//var Residue []string

//	type Diler struct {
//		DilersValue string
//		DilersMasti string
//	}
// type BlackJack struct {
// 	mast  string
// 	cards string
// }

// func nBlackJack(mast string, cards string) BlackJack {
// 	return BlackJack{mast: mast, cards: cards}
// }

//	fmt.Println(deck)}

/////////// игрок и дилер//////////////////

// fmt.Println(DilersValue)

// for j := 0; j <= 4; j++ {
// 	index := rand.Intn(len(deck))
// 	value := deck[index]
// 	PlayerValue = append(PlayerValue, fmt.Sprintf("%s", value))

// }

////остаток колоды/////

// for _,v := range deck{
// 	for _, value := range DilersValue {
// 	fmt.Println(v)
// value

// 	if v != value {
// 		//append()
// 		}
// 	}
// }

///рука
/////      ////////////

//	fmt.Println(PlayerValue)

// 	for len(result) < n {
// 		index := rand.Intn(len(source))
// 		if !usedIndexes[index] {
// 			usedIndexes[index] = true
// 			result = append(result, source[index])
// 		}
// 	}
// 	return result
// }

//rand.Shuffle(len(Cards), func(i, j int) { Cards[i], Cards[j] = Cards[j], Cards[i] })

// for i := 0; i < len(Cards); i++ {
// 	for j := 0; j < len(Mast); j++ {
// 		deck = append(deck, Deck{mast:  Mast, })
// 	}
// }

//for i := 0; i <= 4; i++ {
// 		randomIndex := rand.Intn(len(Cards))
// 		DilersValue = append(DilersValue, Cards[randomIndex])
// 		rand.Shuffle(len(Cards), func(i, j int) { Cards[i], Cards[j] = Cards[j], Cards[i] })

// 	}

// var Masti []string
// for _, v := range Mast {

// 	Masti = append{mast:Masti}
// 	fmt.Println(Masti,)
// }

//Player := make(map[string]string)
// 	unique := make(map[string]bool)

// 	/////////////Cоздание Дилера///////////////////////////////////////////////////////////////////////
// 	var DilersValue []string //////////////значения колоды
// 	for i := 0; i <= 4; i++ {
// 		randomIndex := rand.Intn(len(Cards))
// 		DilersValue = append(DilersValue, Cards[randomIndex])
// 		rand.Shuffle(len(Cards), func(i, j int) { Cards[i], Cards[j] = Cards[j], Cards[i] })

// 	}

// 	for _, v := range DilersValue {
// 		if _, exists := unique[v]; !exists {
// 			unique[v] = true
// 			//		fmt.Println(v)
// 		}

// 	}

// 	//fmt.Println(DilersValue)

// 	var DilersMasti []string ///////масти колоды
// 	for i := 0; i <= 4; i++ {
// 		randomIndex := rand.Intn(len(Mast))
// 		DilersMasti = append(DilersMasti, Mast[randomIndex])

// 	}

// 	//fmt.Println(DilersMasti)

// 	fmt.Println("Diler")
// 	for i := 0; i <= 4; i++ {

// 		fmt.Println(DilersMasti[i]+" "+DilersValue[i]+";", " ")

// 	}

// 	fmt.Println("----------------------------------------------------------------------------------------------------------------------------------------")
// 	////////СОЗДАНИЕ ИГРОКА/////////////
// 	//перемешивание карт колоды по значениям
// 	var CardsValue []string //знаечния колоды
// 	for i := 0; i <= 4; i++ {
// 		randomIndex := rand.Intn(len(Cards))
// 		CardsValue = append(CardsValue, Cards[randomIndex])
// 		rand.Shuffle(len(Cards), func(i, j int) { Cards[i], Cards[j] = Cards[j], Cards[i] })

// 	}

// 	for _, v := range CardsValue {
// 		if _, exists := unique[v]; !exists {
// 			unique[v] = true
// 			//fmt.Println(v)
// 		}
// 	}

// 	// for _ , v := range CardsValue {
// 	// //	if _, exists := DilersValue[v];!exists {

// 	// 	}
// 	// }
// 	// 	masti = append(masti, )
// 	///////////////////////////////////////////////////////////////////////////////////////////////////////

// 	//перемешивание карт колоды по мастям
// 	var masti []string //масти колоды
// 	for i := 0; i <= 4; i++ {
// 		randomIndex := rand.Intn(len(Mast))
// 		masti = append(masti, Mast[randomIndex])

// 	}

// 	//fmt.Println(masti)

// 	//вывести игрока на консоль  PLAYER-1
// 	fmt.Println("Player")
// 	for i := 0; i <= 4; i++ {

// 		fmt.Println(masti[i]+" "+CardsValue[i]+";", " ")

// 	}

// 	fmt.Println() //Player[CardsValue[i]] = masti[i]

// }

/*
	Реализовать поэтапно:
	- Создание колоды;
	- Перемешивание карт колоды;
	- Создать как минимум одного игрока и дилера;
	- Раздать каждому игроку и дилеру по 5 карт;
	- Вывести в консоль карты игроков и дилера в формате `Карта {масть} - {номинал}`
*/

// ВАШ КОД ТУТ...

// func printCards(cards []string) {
// 	for i, val := range cards {
// 		fmt.Printf("Карта #%d - %s\n", i, val)
// 	}
// }
