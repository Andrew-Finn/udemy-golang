//Create your own type. Have the underlying type be an int.
//create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
//in func main
//print out the value of the variable “x”
//print out the type of the variable “x”
//assign 42 to the VARIABLE “x” using the “=” OPERATOR
//print out the value of the variable “x”

package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}