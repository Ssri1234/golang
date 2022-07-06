package main

import "fmt"

type deck []string

func createNewDeck() deck {
	//newDeck := deck{"Ace of Spades", "Two of Diamonds"}
	//newDeck = append(newDeck, "Three of clubs")
	suits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	numbers := []string{"Ace", "two", "three", "four"}

	newDeck := []string{}
	for _, suit := range suits {
		for _, number := range numbers {
			newDeck = append(newDeck, number+" of "+suit)
		}
	}

	return newDeck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//func (d deck) shuffle() {
//d[0], d[15] = d[15], d[0]
//d[1], d[2] = d[3], d[4]
//d[3], d[4] = d[5], d[6]
//d[5], d[6] = d[7], d[8]
//d[7], d[8] = d[9], d[10]
//d[9], d[10] = d[11], d[12]
//d[11], d[12] = d[13], d[14]
//d[13], d[14] = d[0], d[15]

//for i, card := range d {
//fmt.Println(i, card)
//}
//}
