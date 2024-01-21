package variables

import (
	"fmt"
	"time"

	)
	

var Name string
var State bool
var Salary float32
var Today time.Time



func RestOfVariables(){
	Name = "Diego"
	State = true
	Salary = 100.50
	Today = time.Now()

	fmt.Println("Name: ", Name)
	fmt.Println("State: ", State)
	fmt.Println("Salary: ", Salary)
	fmt.Println("Today: ", Today)
}