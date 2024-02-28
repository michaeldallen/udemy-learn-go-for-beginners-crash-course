package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)

	userName := readString("What is your name?")

	age := readInt("How old are you?")
	fmt.Printf("Your name is %s. You are %d years old\n", userName, age)

}

func prompt() {
	fmt.Print("-> ")

}

func readString(msg string) string {

	for {
		fmt.Println(msg)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "" {
			fmt.Println("type something")
		} else {
			return userInput
		}

	}

}

func readInt(msg string) int {

	for {

		fmt.Println(msg)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}

}
