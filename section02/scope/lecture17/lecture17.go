package lecture17

import (
	"fmt"
)

var PackageVar string = "PackageVar - package-level variable for package 'section02'"

func PrintMe(a string, b string) {
	fmt.Println(a, b, PackageVar)

}
