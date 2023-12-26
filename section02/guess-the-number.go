package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt string = "and don't type your number in, just  press ENTER when ready."

func answer(first int, second int, third int) int {
	return (first * second) - third
}

func play(firstNumber int, secondNumber int, subtraction int) {
	// sed
	rand.Seed(time.Now().UnixNano())
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
	fmt.Println("The answer is", answer(firstNumber, secondNumber, subtraction))

}

func main() {

	var firstNumber int = rand.Intn(8) + 2
	var secondNumber int = rand.Intn(8) + 2
	var subtraction int = rand.Intn(8) + 2

	play(firstNumber, secondNumber, subtraction)
}
