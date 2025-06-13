package game

type HandStrength int

// List of all combinations
const (
	HighCard HandStrength = iota
	Pair
	TwoPairs
	Set
	Straight
	Flush
	FullHouse
	Quads
	StraightFlash
	RoyalFlush
)
