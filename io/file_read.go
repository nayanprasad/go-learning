package main

import (
	// "fmt"
	"io"
	"os"
)

func main() {

	args := os.Args
	fileName := args[1]

	file, err := os.Open(fileName)

	if err != nil {
		os.Exit(1)
	}

	// data := make([]byte, 100)

	// _, _ = file.Read(data)

	// fmt.Println(string(data))

	io.Copy(os.Stdout, file)

}
