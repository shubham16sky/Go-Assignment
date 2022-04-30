package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var number int
	number = rand.Int()
	fmt.Printf("Randomly generated integer is %v", number)

}
