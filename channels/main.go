package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.co",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)  // create channel type

	for _, link := range links {
		go checkLinks(link, c)
	}
	for i := 0; i < len(links); i++ {   // Receiving the respone depending upon channnel length
		fmt.Println(<-c)
	}

}

func checkLinks(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- "Might be down"
		return
	}
	fmt.Println(link, "is up!")
	c <- "Is up"
}
