package files

import (
	"bufio"
	"fmt"

	"os"

	"github.com/DiegoNGonzalez/goDesde0/ejercicios"
)

var filename string= "./files/txt/tabla.txt"

func GrabaTabla(){
	var texto string = ejercicios.Multiplicador()
	archivo, err := os.Create(filename)
	if err != nil{
		fmt.Println("Error al crear al archivo"+ err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla(){
	var texto string = ejercicios.Multiplicador()
	if !Append(filename, texto){
		fmt.Println("Error al concatenar contenido.")
	}
}

func Append(filen string, texto string) bool{
	arch, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append"+err.Error())
		return false
	}

	_, err= arch.WriteString(texto)
	if err != nil{
		fmt.Println("Error durante el Writestring"+err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo(){
	archivo,err:= os.Open(filename)
	if err != nil{
		fmt.Println("Error en lectura"+err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan(){
		registro := scanner.Text()
		fmt.Println(">"+registro)
	}
	archivo.Close()
}