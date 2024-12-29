package main

import "fmt"

func main() {
	card := newDeckFromFile("my_cards.txt")
	fmt.Println("my cards are ", card)
	card.print()
}
