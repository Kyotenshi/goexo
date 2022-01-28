package main

import (
	"fmt"
)

func main() {
	name := "Johan"
	//CODE
	switch {
	case name == "Johan":
		fmt.Println("Good guess, my name is indeed: ", name)
	case len(name) == len("Johan"):
		fmt.Println("Nope, it's not my name, but we have the same number of letters !", len(name))
	default:
		fmt.Println("Unlucky, we have nothing in common")

	}
}
