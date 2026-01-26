package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a number: ")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())

	if num%2 == 0 {
		fmt.Printf("Num %v is odd", num)
	} else {
		fmt.Printf("Num %v is even", num)
	}

}
