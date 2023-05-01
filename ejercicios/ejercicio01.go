package ejercicios

import (
	"fmt"
	"strconv"
)

func ConviertoAInt(cadena string)(int,string){

	numero, _:= strconv.Atoi(cadena)
	var mensaje string
	if numero >100{
		mensaje ="Es mayor a 100"
	}else{
		mensaje="Es menor a 100"
	}
	fmt.Println(numero)
	fmt.Println(mensaje)
	return numero, mensaje
}

func ConvNumerico(texto string)(int,string){
	num, err :=strconv.Atoi(texto)
	if err != nil{
		return 0, "Hubo un error"+err.Error()
	}
	if num >100{
		return num, "Es mayor a 100"
	}else {
		return num, "Es menor a 100"
	}
}