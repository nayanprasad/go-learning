package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// fmt.Println(resp)

	// body := []byte{}
	body := make([]byte, 99999)
	// why 99999? because the body is a stream of data and we need to read it all, if we don't know the size of the body, we can use make([]byte, 0) and read it until the end of the stream
	// but in this case we know the size of the body, so we can use make([]byte, 99999)
	// so we can read the body until the end of the stream
	// otherwise we will get a partial body

	resp.Body.Read(body)

	fmt.Println(string(body))
}
