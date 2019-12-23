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

	for _, link := range links {
		checkLinks(link)
	}
}

func checkLinks(link string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		return
	}
	fmt.Println(link, "is up!")
}
