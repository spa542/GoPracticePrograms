package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	// This is just to show that there is a function that can read the exact number of bytes from the response body
	// body, err := io.ReadAll(resp.Body)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
