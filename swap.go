package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Printf("Enter the two numbers A and B :")
	fmt.Scanf("%v%v", &a, &b)
	a, b = swap(a, b)
	//fmt.Println("After reversing\n A and B is " + swap(a, b))
	fmt.Printf("After swapping Value of A and B is \n A = %v \n B = %v \n", a, b)
}

func swap(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b

	return a, b
}
