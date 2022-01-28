package main

import (
	"fmt"
)

func main() {
	name := "Johny"
	//CODE
	if name == "Johan" {
		fmt.Println("It's indeed my name :", name)
	} else if len(name) == len("Johan") {
		fmt.Println("It's not my name, but fun fact, we have the same number of letters !", len(name))
	} else {
		println("Unlucky, we have nothing in common")
	}
}
