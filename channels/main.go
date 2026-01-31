package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Create a channel
	// Make is a function that will create variable of some type
	c := make(chan string)

	// We do not want to make these requests one by one. Instead, lets do them concurrently using Go routines!
	for _, link := range links {
		go checkLink(link, c)
	}

	// Receive the value from the channel
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// Receive a number of messages equal to the number of requests that we make
	// for range links {
	// for {
	// 	// fmt.Println(<-c)
	// 	go checkLink(<-c, c)
	// }
	// Alternative syntax - equivalent to above
	for l := range c {
		// time.Sleep(5 * time.Second)
		// go checkLink(l, c)
		// Function Literal or lambda / anonymous function
		go func(link string) {
			time.Sleep(5 * time.Second)
			// The variable l is being referenced from the outer scope
			// l could be changed in memory before the request is actually resolved, therefore, we pass l as an argument to the anonymous function
			checkLink(link, c)
		}(l) // This is to evoke the function right after definition
	}
}

// Even here we have to declare what type the channel is planning to pass
func checkLink(link string, c chan string) {
	// time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!")
		// c <- "Might be down I think"
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	// c <- "Yep its up"
	c <- link
}
