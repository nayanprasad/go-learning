package main

import (
	"fmt"
	"net/http"
	"sync"
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

	pingLinksSeqentially(links)

	fmt.Println("==============")

	pingLinksWithGoRoutines(links)

}

func pingLinksWithGoRoutines(links []string) {

	start := time.Now()

	var wg sync.WaitGroup //

	for _, link := range links {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			pingLink(l)
		}(link)
		// go pingLink(link)
	}

	wg.Wait()

	/*
		conceptually, a WaitGroup is a counter plus a blocker.

		Add(n) - increments the counter

		Done() - decrements the counter by 1

		Wait() - blocks until the counter reaches 0
	*/

	elapsed := time.Since(start)
	fmt.Println("took ", elapsed)

}

func pingLinksSeqentially(links []string) {

	start := time.Now()

	for _, link := range links {
		pingLink(link)
	}

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("took ", elapsed)
}

func pingLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("down ", link)
		return
	}

	fmt.Println("up ", link)
}
