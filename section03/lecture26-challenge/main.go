package main

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
	HasCat         bool
	HasDog         bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavoriteNumber = readFloat("What is your favorite Number?")

	fmt.Printf("Your name is %s, you are %d years old and your favorite number is %.2f.\n",
		user.UserName,
		user.Age,
		user.FavoriteNumber)

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

func readFloat(msg string) float64 {

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

func readBool(msg string) bool {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {

		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'n' || char == 'N' {
			return false
		}
		if char == 'y' || char == 'Y' {
			return true
		}

	}
}
