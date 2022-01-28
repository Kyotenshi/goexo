package main

import "fmt"

func main() {
	fmt.Println(stringToBool("omfg"))
}

func stringToBool(conversion string) bool {
	if conversion == "true" {
		return true
	} else {
		return false
	}
}
