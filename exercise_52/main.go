package main

import (
	"fmt"
)

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	delete(colors, "white")
	fmt.Println(colors)

	printMap(colors)

	var colors2 map[string]string
	colors2 = make(map[string]string)
	colors2["white"] = "#ffffff"
	fmt.Println(colors2)

	colors3 := make(map[string]string)
	colors3["white"] = "#ffffff"
	fmt.Println(colors3)
}
