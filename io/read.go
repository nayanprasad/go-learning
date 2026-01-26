package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	bs, err := ioutil.ReadFile("file d")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := string(bs)

	fmt.Println(s)
}
