//Program to show the different ways of writing functions

package main

import (
	"fmt"
)

func main() {

	var a int
	var b int

	fmt.Printf("Enter the A and B : ")
	fmt.Scanf("%v%v", &a, &b)

	moreThanOneParameter(a, b)

	var c, d int

	c, d = diffReturnValues(a, b)

	fmt.Printf("variables returned by diffReturnValues function : %v , %v", c, d)

}

func moreThanOneParameter(a, b int) {
	fmt.Printf("A and B is %v %v\n\n", a, b)
}

func diffReturnValues(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}
