package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	words := []string{"first", "second", "third", "fourth", "fifth"}

	for index, word := range words {
		go printSomething(fmt.Sprintf("This is the %d %s thing to be printed.", index+1, word))
	}

	time.Sleep(1 * time.Second)
	printSomething("This is the second thing to be printed.")
}
