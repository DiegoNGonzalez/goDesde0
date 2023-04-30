package main

import (
	"fmt"

	"github.com/DiegoNGonzalez/goDesde0/variables"
)

func main(){

	estado,texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

}