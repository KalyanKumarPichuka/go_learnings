package main

import "fmt"

func main() {
	card := newDeck()
	card.print()
	card.shuffle()
	fmt.Println("Shuffled deck")
	card.print()
}
