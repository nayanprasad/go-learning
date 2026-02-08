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
		"http://instagram.com",
		"http://x.com",
		"http://nayanprasad.com",
	}

	// pingLinksWithChannel(links)

	statusChecker(links)

}

func statusChecker(links []string) {
	c := make(chan string)

	for _, link := range links {
		go pingLink(link, c)
	}

	for link := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			pingLink(l, c)
		}(link)
	}
}

func pingLinksWithChannel(links []string) {
	c := make(chan string)

	for _, link := range links {
		go pingLink(link, c)
	}

	for link := range c { // this run forever
		fmt.Println(link)
	}
}

func pingLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Down -> ", link)
		c <- link
		return
	}

	fmt.Println("Up -> ", link)
	c <- link
}
