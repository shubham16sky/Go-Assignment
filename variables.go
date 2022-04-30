//program to show the different ways of declaring variables in different scopes

package main

import "fmt"

var globalVar = 10

func main() {

	var localVar = 20

	fmt.Printf("Global variable is : %v and Local Variable is : %v\n", globalVar, localVar)

	if ifVar := 30; ifVar > 50 {
		fmt.Printf("ifVar is greater than 50")
	}
	fmt.Printf("ifVar is smaller than 50")

}
