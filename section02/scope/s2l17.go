package main

import (
	"section02/lecture17"
)

var myVar string = "myVar - package-level variable for package 'main'"

func main() {

	var blockVar string = "blockVar - block-level variable"

	// variables

	// declare a package-level variable for the main pacakge named "myVar"

	// declare a block-level variable for the main function called "blockVar"

	// declare a package-level variable named "PackageVar"

	// create an exported function called PrintMe

	// in main print out the values of
	// * myVar
	// * blockVar
	// * packageVar
	// on one line
	// using the PrintMe function

	lecture17.PrintMe(myVar, blockVar)

}
