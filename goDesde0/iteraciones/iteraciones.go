package iteraciones

import "fmt"


func Iterar(){

	for i:=0;i<10;i++{
		fmt.Println(i)
	}

	for i:=0;i<100;i+=5{
		
		fmt.Println(i)
	}

	for i:=0;i<10;i++{
		if i==5{
			fmt.Println("Cinco")
			continue
		}
		fmt.Println(i)
	}
}