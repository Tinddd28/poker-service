package game

import (
	"math/rand"
	"time"
)

// Representation of a pack of cards in the game
type Deck struct {
	Cards  []Card
	Length int
}

// Representation of a card in the game with suit and rank
type Card struct {
	Rank Rank
	Suit Suit
}

// Return Deck with all rank cards 2-10, J, Q, K, A and suits: Hearts, Diamonds, Clubs, Spades
func NewDeck() *Deck {
	cards := make([]Card, 0, CountCards)

	for _, suit := range AllSuits {
		for _, rank := range AllRanks {
			cards = append(cards,
				Card{
					Rank: rank,
					Suit: suit,
				},
			)
		}
	}

	return &Deck{
		Cards:  cards,
		Length: len(cards),
	}
}

// ShuffleDeck shuffles the deck of cards using the Fisher-Yates algorithm
// #NOTE: In the future: should rework
func (d *Deck) ShuffleDeck() {
	s := rand.NewSource(time.Now().Unix())
	newRand := rand.New(s)
	for i := len(d.Cards) - 1; i > 0; i-- {
		j := newRand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func (d *Deck) GetCard() Card {
	card := d.Cards[d.Length-1]
	d.Cards = d.Cards[:d.Length-1]
	d.Length--
	return card
}
