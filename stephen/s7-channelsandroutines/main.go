package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	fmt.Println("go routine added")
	_, err := http.Get(link)
	if err != nil {
		println(link, "might be down!")
		return
	}
	println(link, "is up!")
}
