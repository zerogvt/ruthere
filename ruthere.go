package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	targets := []string{
		"http://google.com",
		"http://bing.com",
		"http://duckduckgo.com",
		"http://apple.com",
	}

	c := make(chan string)
	for _, t := range targets {
		go ping(t, c)
	}
	for msg := range c {
		go func(url string) {
			time.Sleep(5 * time.Second)
			ping(url, c)
		}(msg)
	}
}

func ping(url string, c chan string) bool {
	_, err := http.Get(url)
	res := err == nil
	if res {
		log.Println(url, "is reachabe")
	} else {
		log.Println(url, "is not reachabe")
	}
	c <- url
	return res
}
