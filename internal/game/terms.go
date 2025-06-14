package game

// Colors of the suits in a standard deck of playing cards for CLI
const (
	Reset = "\033[0m"  // Return to default color
	Red   = "\033[31m" // Red color
	Black = "\033[30m" // Black color
)

// Suit
const (
	Spades      = Black + "\xE2\x99\xA0" + Reset
	SpadesStr   = "spades"
	Clubs       = Black + "\xE2\x99\xA3" + Reset
	ClubsStr    = "clubs"
	Hearts      = Red + "\xE2\x99\xA5" + Reset
	HeartsStr   = "hearts"
	Diamonds    = Red + "\xE2\x99\xA6" + Reset
	DiamondsStr = "diamonds"
)

// Suit represent the suit of a card in a standard deck of playing cards
type Suit uint8

const (
	SuitSpades   Suit = iota // spades suit
	SuitClubs                // clubs suit
	SuitHearts               // hearts suit
	SuitDiamonds             // diamonds suit
)

// Map for converting suit to its string representation
var suitToStringMap = map[Suit]string{
	SuitSpades:   SpadesStr,
	SuitHearts:   HeartsStr,
	SuitClubs:    ClubsStr,
	SuitDiamonds: DiamondsStr,
}

// Map for converting suit to its Unicode symbol with color
var suitToSymbolMap = map[Suit]string{
	SuitSpades:   Spades,
	SuitHearts:   Hearts,
	SuitClubs:    Clubs,
	SuitDiamonds: Diamonds,
}

// String returns the string representation of the Suit
func (s Suit) String() string {
	if name, ok := suitToStringMap[s]; ok {
		return name
	}
	return "unknown"
}

// Symbol returns the Unicode symbol with color for the suit
func (s Suit) Symbol() string {
	if sym, ok := suitToSymbolMap[s]; ok {
		return sym
	}
	return "?"
}

// Rank represents the rank of a card in a standard deck of playing cards
// Ranks acceptable values are from 2 to 14, where:
// 2 - 10 represent the numbered cards,
// 11 represents Jack, 12 represents Queen, 13 represents King and 14 represents Ace
type Rank uint8

const (
	RankTwo   Rank = iota + 2 // 2
	RankThree                 // 3
	RankFour                  // 4
	RankFive                  // 5
	RankSix                   // 6
	RankSeven                 // 7
	RankEight                 // 8
	RankNine                  // 9
	RankTen                   // 10
	RankJack                  // 11
	RankQueen                 // 12
	RankKing                  // 13
	RankAce                   // 14
)

// Nominal map for converting Rank to its string representation
var nominalMap = map[Rank]string{
	RankTwo:   "2",
	RankThree: "3",
	RankFour:  "4",
	RankFive:  "5",
	RankSix:   "6",
	RankSeven: "7",
	RankEight: "8",
	RankNine:  "9",
	RankTen:   "10",
	RankJack:  "J",
	RankQueen: "Q",
	RankKing:  "K",
	RankAce:   "A",
}

// String returns the string representation of the Rank
func (r Rank) String() string {
	if name, ok := nominalMap[r]; ok {
		return name
	}

	return "unknown"
}

var AllSuits = []Suit{SuitSpades, SuitHearts, SuitDiamonds, SuitClubs}

var AllRanks = []Rank{RankTwo, RankThree, RankFour, RankFive, RankSix, RankSeven, RankEight,
	RankNine, RankTen, RankJack, RankQueen, RankKing, RankAce}

// Counts for various game elements
const (
	AllCards          = 52 // Total cards on a standard deck
	CountCards        = 2  // Number of cards in a player`s hand
	CountSuit         = 4  // Number of suits in a standard deck
	CountFlop         = 3  // Number of flop cards on the table
	CountNominals     = 13 // Number of a nominal cards in a standard deck
	CountCombinations = 21 // Number of combinations  C(7, 5) = 7! / (5! * (7 - 5)!) = 21
	MaxCard           = 5
)
