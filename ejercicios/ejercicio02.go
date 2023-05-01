package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func Multiplicador(){

	scanner := bufio.NewScanner(os.Stdin)
	/* fmt.Println("Ingrese un numero: ")
	if scanner.Scan(){
		numero, err = strconv.Atoi(scanner.Text())
		for err!=nil{
			fmt.Println("error ingrese un numero: ")
			if scanner.Scan(){
				numero, err = strconv.Atoi(scanner.Text())
			}
		}
	} */
	for{
		fmt.Println("Ingrese un numero: ")
		if scanner.Scan(){
			numero,err = strconv.Atoi(scanner.Text())
			if err != nil{
				continue
			}else{
				break
			}
		}
	}

	for i:=1; i<=10; i++{
		fmt.Printf("%d x %d = %d \n", numero,i,i*numero)
		fmt.Println(numero*i)
	}
}