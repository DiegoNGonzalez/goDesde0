package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*fmt.Println("Hola Mundo")
	variables.ShowInt()
	variables.RestOfVariables()
	state,text := variables.ConverToText(35)
	fmt.Println("State: ", state)
	fmt.Println("Text convert: ", text)*/

	if os := runtime.GOOS; os == "Linux." || os == "Os X. " {
		fmt.Println("This doesn't is Windows")

	} else {
		fmt.Println("This is Windows")
	}

	switch os := runtime.GOOS; os {
	case "Linux":
		fmt.Println("This is Linux")
	case "Os X":
		fmt.Println("This is Os X")
	default:
		fmt.Printf("%s. \n", os)
	}
}
