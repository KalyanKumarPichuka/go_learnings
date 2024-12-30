package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	words := []string{"first", "second", "third", "fourth", "fifth"}

	wg.Add(len(words))
	for index, word := range words {
		go printSomething(fmt.Sprintf("This is the %d %s thing to be printed.", index+1, word), &wg)
	}
	wg.Wait()
}
