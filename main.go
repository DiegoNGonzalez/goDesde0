package main

import (
	"fmt"
	"runtime"
)

func main(){

	/*estado,texto := variables.ConviertoATexto(15700)
	fmt.Println(estado)
	fmt.Println(texto)
*/
	if os:= runtime.GOOS; os=="Linux." || os== "Os X."{
		fmt.Println("Esto no es Windows, es ", os)

	}else {
		fmt.Println("Esto es Windows")
	}

	switch os:= runtime.GOOS; os{
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	} 
}