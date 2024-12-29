package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	fmt.Println(resp)
	fmt.Println("Status code:", resp.Status)

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
