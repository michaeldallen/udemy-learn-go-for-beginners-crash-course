package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {
			fmt.Println(userInput)
		}
	}
}
