package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
	}

	// Create a new channel
	channel := make(chan string)

	for _, link := range links {
		// go keyword is create a new child routine.
		// main routine doesn't wait the child routine finished
		go checkLink(link, channel)
	}

	// Refetch link from channel
	for l := range channel {
		// This is an anonymous function (called function literal in Go)
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, channel)
		}(l)
	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		// Assign into channel
		channel <- link
		return
	}
	fmt.Println(link, "is up")
	// Assign into channel
	channel <- link
}
