package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName       string
	Age            int
	FavoriteNumber float
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.userName = readString("What is your name?")
	user.age = readInt("How old are you?")
	user.favoriteNumber = readFloat("What is your favorite Number?")

	fmt.Printf("Your name is %s, you are %d years old and your favorite number is %.2f.\n",
		user.userName,
		user.age,
		user.favoriteNumber)

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

func readInt(msg string) int {

	for {

		fmt.Println(msg)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}

}
