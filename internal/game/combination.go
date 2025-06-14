package game

import mathutils "pokerGo/pkg/mathUtils"

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

// Evaluate(handCards + tableCards)
func getCombinations(cards []Card, k int) [][]Card {
	n := len(cards)
	if k < 0 || k > n || k != 5 {
		return nil
	}
	combs := mathutils.FactorialRecursive(n) / (mathutils.FactorialRecursive(n-k) * mathutils.FactorialRecursive(k))

	combinations := make([][]Card, 0, combs)
	indices := make([]int, k)

	for i := range indices {
		indices[i] = i
	}

	for {
		cur := make([]Card, k)
		for i, index := range indices {
			cur[i] = cards[index]
		}

		combinations = append(combinations, cur)

		i := k - 1
		for i >= 0 && indices[i] == n-k+i {
			i--
		}

		if i < 0 {
			break
		}

		indices[i]++

		for j := i + 1; j < k; j++ {
			indices[j] = indices[j-1] + 1
		}
	}

	return combinations
}

func Check(cards []Card, player Player) HandStrength {
	combinations := getCombinations(cards, MaxCard)
	for _, comb := range combinations {
		if isHighCard(comb) != nil {
			player.Strength = HighCard
		}
		if isPair(comb) != nil {
			player.Strength = Pair
		}
		if isTwoPairs(comb) != nil {
			player.Strength = TwoPairs
		}
		if isSet(comb) != nil {
			player.Strength = Set
		}
		if isStraight(comb) != nil {
			player.Strength = Straight
		}
		if isFlush(comb) != nil {
			player.Strength = Flush
		}
		if isFullHouse(comb) != nil {
			player.Strength = FullHouse
		}
		if isQuads(comb) != nil {
			player.Strength = Quads
		}
		if isStraightFlush(comb) != nil {
			player.Strength = StraightFlash
		}
		if isRoyalFlush(comb) != nil {
			player.Strength = RoyalFlush
		}
	}
	return 1
}

func isPair(cards []Card) []Card {
	return nil
}

func isTwoPairs(cards []Card) []Card {
	return nil
}

func isSet(cards []Card) []Card {
	return nil
}

func isStraight(cards []Card) []Card {
	return nil
}

func isFlush(cards []Card) []Card {
	return nil
}

func isFullHouse(cards []Card) []Card {
	return nil
}

func isQuads(cards []Card) []Card {
	return nil
}

func isStraightFlush(cards []Card) []Card {
	return nil
}

func isRoyalFlush(cards []Card) []Card {
	return nil
}

func isHighCard(cards []Card) []Card {
	return nil
}
