package main

import (
	"fmt"
	"strings"
)

func main() {

	greeting := "hi there!!!"

	fmt.Println(greeting)

	greetingBytes := []byte(greeting)

	fmt.Println(greetingBytes)

	greetingsSlice := []string{"hi there", "bi there", "ci there", "ai there"}

	fmt.Println(greetingsSlice)

	greetingsString := strings.Join(greetingsSlice, ",")

	fmt.Println(greetingsString)

	fmt.Println([]byte(greetingsString))

}
