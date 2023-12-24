package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt string = "and press ENTER when ready."

func main() {

	var firstNumber int = 2
	var secondNumber int = 5
	var subtraction int = 7
	//	var answer int

	reader := bufio.NewReader(os.Stdin)

	// display welcome/instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take user through the game
	fmt.Println("Multiply your number by ", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give user the answer

}
