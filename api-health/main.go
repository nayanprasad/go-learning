package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hi"))
}

func main() {
	http.HandleFunc("/health", healthHandler)

	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
}
