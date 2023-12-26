package lecture16

import (
	"fmt"
)


func notExported() {
	fmt.Println("notExported")

}

func Exported() {
	fmt.Println("Exported")

}
