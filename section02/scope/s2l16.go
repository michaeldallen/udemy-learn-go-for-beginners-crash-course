package main

import (
	"fmt"
	"section02/lecture16"
)

var one string = "one"

func main() {

	var one string = "this is a block-level variable"

	fmt.Println(one)

	myFunc(one)

	fmt.Println("from packageone:", lecture16.PublicVar)

	lecture16.Exported()

}

func myFunc(one string) {
	//	var one string = "the number one"
	fmt.Println(one)
}
