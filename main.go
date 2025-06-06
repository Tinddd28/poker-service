package main

import (
	"fmt"
	"math/rand"
	"time"
)

// colors
const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Black = "\033[30m"
)

// suit
const (
	spades   = Black + "\xE2\x99\xA0" + Reset
	clubs    = Black + "\xE2\x99\xA3" + Reset
	hearts   = Red + "\xE2\x99\xA5" + Reset
	diamonds = Red + "\xE2\x99\xA6" + Reset
)

// nominal
var nominalMap = map[int]string{
	1:  "A",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "10",
	11: "J",
	12: "Q",
	13: "K",
}

// counts
const (
	allCards      = 52
	countCards    = 2
	countSuit     = 4
	countFlop     = 4
	countNominals = 13
)

type Deck struct {
	Cards []Card
}

type Card struct {
	NominalInt int
	NominalStr string
	Suit       string
}

func NewDeck() *Deck {
	cards := make([]Card, 0)
	for i := 1; i < countNominals+1; i++ {
		cards = append(cards, Card{
			NominalInt: i,
			NominalStr: nominalMap[i],
			Suit:       spades,
		})
		cards = append(cards, Card{
			NominalInt: i,
			NominalStr: nominalMap[i],
			Suit:       hearts,
		})
		cards = append(cards, Card{
			NominalInt: i,
			NominalStr: nominalMap[i],
			Suit:       clubs,
		})
		cards = append(cards, Card{
			NominalInt: i,
			NominalStr: nominalMap[i],
			Suit:       diamonds,
		})
	}
	return &Deck{
		Cards: cards,
	}
}

func (d *Deck) ShuffleDeck() {
	s := rand.NewSource(time.Now().Unix())
	newRand := rand.New(s)
	for i := len(d.Cards) - 1; i > 0; i-- {
		j := newRand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func main() {
	deck := NewDeck()
	for i := 0; i < 3; i++ {
		deck.ShuffleDeck()
	}
	var count int
	fmt.Printf("input players count... -> ")
	fmt.Scan(&count)
	count = 3
	// for _, d := range deck.Cards {
	// 	for k, v := range d {
	// 		fmt.Println(nominalMap[v], k)
	// 	}
	// }
	// var players [][]Card
	players := make([][]Card, count)
	cardMinus := 1
	for i := 0; i < count; i++ {
		player := make([]Card, countCards)
		for j := 0; j < countCards; j++ {
			player[j] = deck.Cards[len(deck.Cards)-cardMinus]
			cardMinus++
		}
		players[i] = player
	}
	fmt.Println("first player cards")
	for i, p := range players {
		fmt.Printf("%d cards: ", i+1)
		for _, c := range p {
			fmt.Printf("%s %s ", c.NominalStr, c.Suit)
		}
		fmt.Println()
	}
}
