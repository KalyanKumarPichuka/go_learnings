package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}

func main() {

	links := []string{"hhtp://google.com", "http://facebook.com", "http://golang.org", "http://amazon.com"}

	for _, link := range links {
		checkLink(link)
	}
}
