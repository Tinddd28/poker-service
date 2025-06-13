package game

type Table struct {
	Deck       *Deck
	Players    []Player
	StashCards []Card
	Street     []Card
}
type Player struct {
	Hand [2]Card
	Name string
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
		for pl := range t.Players {
			t.Players[pl].Hand[i] = t.Deck.GetCard()
		}
	}
}

func (t *Table) Flop() {
	for i := 0; i < CountFlop; i++ {
		t.Street = append(t.Street, t.Deck.Cards[t.Deck.Length-1])
		t.Deck.Length--
	}
}
