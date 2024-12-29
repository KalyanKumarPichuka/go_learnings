package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	for i := 1; i <= 10; i++ {
		if isEven(i) {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}
}
