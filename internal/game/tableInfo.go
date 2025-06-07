package game

import (
	"math/rand"
	"time"
)

type Table struct {
	Deck       *Deck
	Players    []Player
	StashCards []Card
	Street     []Card
}

func CreateTable() *Table {
	table := Table{}
	deck := NewDeck()
	for i := 0; i < 3; i++ {
		deck.ShuffleDeck()
	}
	table.Deck = deck
	return &table
}

func (t *Table) AddPlayer(pl Player) {
	t.Players = append(t.Players, pl)
}

func (t *Table) StartGame() {
	for i := 0; i < CountCards; i++ {
		for pl, _ := range t.Players {
			t.Players[pl].Hand[i] = t.GetCard()
		}
	}
}

func (t *Table) Flop() {
	for i := 0; i < CountFlop; i++ {
		t.Street = append(t.Street, t.Deck.Cards[t.Deck.Length-1])
		t.Deck.Length--
	}
}

type Player struct {
	Hand [2]Card
	Name string
	Bank int
}

type Deck struct {
	Cards  []Card
	Length int
}

type Card struct {
	NominalInt int
	NominalStr string
	Suit       string
	SuitName   string
}

func NewDeck() *Deck {
	cards := make([]Card, 0)
	for i := 1; i < CountNominals+1; i++ {
		cards = append(cards, Card{
			NominalInt: i,
			NominalStr: NominalMap[i],
			Suit:       Spades,
			SuitName:   SpadesStr,
		})
		cards = append(cards, Card{
			NominalInt: i,
			NominalStr: NominalMap[i],
			Suit:       Hearts,
			SuitName:   HeartsStr,
		})
		cards = append(cards, Card{
			NominalInt: i,
			NominalStr: NominalMap[i],
			Suit:       Clubs,
			SuitName:   ClubsStr,
		})
		cards = append(cards, Card{
			NominalInt: i,
			NominalStr: NominalMap[i],
			Suit:       Diamonds,
			SuitName:   DiamondsStr,
		})
	}
	return &Deck{
		Cards:  cards,
		Length: AllCards,
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

func (t *Table) GetCard() Card {
	card := t.Deck.Cards[t.Deck.Length-1]
	t.Deck.Cards = t.Deck.Cards[:t.Deck.Length-1]
	t.Deck.Length--
	return card
}
