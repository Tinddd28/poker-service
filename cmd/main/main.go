package main

import (
	"fmt"
	mathutils "pokerGo/pkg/mathUtils"
)

func main() {

	fmt.Println(mathutils.FactorialRecursive(7) / (mathutils.FactorialRecursive(7-5) * mathutils.FactorialRecursive(5)))

	// deck := game.NewDeck()
	// for i := 0; i < 3; i++ {
	// 	deck.ShuffleDeck()
	// }

	// table1 := game.CreateTable()

	// var count int
	// fmt.Printf("input players count... -> ")
	// fmt.Scan(&count)
	// if (game.AllCards-5)/count < 2.0 {
	// 	fmt.Println("cant create a table!")
	// 	return
	// }

	// for i := 0; i < count; i++ {
	// 	table1.Players = append(table1.Players, game.Player{
	// 		Name: fmt.Sprintf("player %d", i+1),
	// 	})
	// }
	// table1.StartGame()
	// table1.Flop()

	// for i := 0; i < count; i++ {
	// 	fmt.Print(table1.Players[i].Name, " cards: ")
	// 	for _, c := range table1.Players[i].Hand {
	// 		fmt.Printf("%s %s ", c.Rank, c.Suit)
	// 	}
	// 	fmt.Println()

	// 	fmt.Println()
	// }

	// for i := 0; i < game.CountFlop; i++ {
	// 	fmt.Printf("%s %s ", table1.Street[i].Rank, table1.Street[i].Suit)
	// }

	// fmt.Println()
}
