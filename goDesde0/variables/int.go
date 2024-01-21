package variables

import "fmt"

func ShowInt(){
	var intCom int 
	intOf32 := int32(10)
	intOf64 := int64(100)
	intCom = 1

	fmt.Println("intCom: ", intCom)
	fmt.Println("intOf32: ", intOf32)
	fmt.Println("intOf64: ", intOf64)
}