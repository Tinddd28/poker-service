package game

// colors
const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Black = "\033[30m"
)

// suit
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

// nominal
var NominalMap = map[int]string{
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
	AllCards      = 52
	CountCards    = 2
	CountSuit     = 4
	CountFlop     = 3
	CountNominals = 13
)
