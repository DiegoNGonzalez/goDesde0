package main

import (
	"fmt" 
	"github.com/diegoNGonzalez/goDesde0/variables"
)
func main(){
	fmt.Println("Hola Mundo")
	variables.ShowInt()
	variables.RestOfVariables()
	state,text := variables.ConverToText(35)
	fmt.Println("State: ", state)
	fmt.Println("Text convert: ", text)
}