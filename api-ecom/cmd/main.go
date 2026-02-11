package main

import (
	"log"
	"os"
)

func main() {
	cnf := config{
		addr: ":5055",
		db:   dbConfig{},
	}

	api := appplication{
		config: cnf,
	}

	h := api.mount()
	error := api.run(h)
	if error != nil {
		log.Printf("Error: starting the server. %s", error)
		os.Exit(1)
	}
}
