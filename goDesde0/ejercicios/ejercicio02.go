package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Exercise02() {
	scanner := bufio.NewScanner(os.Stdin)
	var number int
	var err error

	for {
		fmt.Println("Ingrese un numero: ")
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}

		for i := 1; i <= 10; i++ {
			fmt.Printf("%d x %d = %d \n", number, i, number*i)
		}

	}
}
