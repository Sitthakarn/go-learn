package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://amazon.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://google.com",
	}

	c := make(chan string)

	for _, link := range links {
		// fmt.Println("Firing go routine for", link)
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

	// for l := range c {
	// 	go checkLink(l, c)
	// }
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
	}
	fmt.Println(link, "link is up!")
	c <- link
}
