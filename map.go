package main

import (
	"fmt"
)

func main() {

	// var users map[int]string
	var users = make(map[int]string)

	// users := map[int]string{
	// 	1: "a",
	// 	2: "b",
	// 	3: "c",
	// }

	users[4] = "d"

	delete(users, 4)

	printMap(users)
}

func printMap(m map[int]string) {
	for key, val := range m {
		fmt.Println(key, val)
	}
}
