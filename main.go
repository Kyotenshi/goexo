package main

import "fmt"

func main() {
	bool, err := fmt.Println(stringToBool("omfg"))

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(bool)
	}
}

func stringToBool(conversion string) (bool, error) {
	if conversion == "true" {
		return true, nil
	} else if conversion == "false" {
		return false, nil
	} else {
		return false, fmt.Errorf("Cannot convert string to bool")
	}
}
