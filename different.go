//Program to show difference between Println() and Printf() function

package main

import "fmt"

func main() {
	var lnVar int = 50
	var fVar int = 100

	//In Printf function we have to use %v to format the output string
	//In println function we can just pass the variable to print it or we can just concatenate them using '+'

	fmt.Println("Variable Printing using Println : ", lnVar)
	fmt.Printf("Variable Printing using Printf :%v", fVar)

}
