package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	s := "hi there"

	err := ioutil.WriteFile("filed", []byte(s), 0666)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
