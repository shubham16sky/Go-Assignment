package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var random int
	random = rand.Intn(500)

	fmt.Println("Randomly generated number is : ", random)
}
