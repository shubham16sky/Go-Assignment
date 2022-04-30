//Program to show creation and use of struct
//Pranit Kumar , USN : 1VA19IS039

package main

import "fmt"

type Dogs struct {
	Name   string
	Color  string
	Age    int
	Weight int
	Breed  string
}

func main() {

	dog1 := Dogs{"Tommy", "Black", 4, 25, "Pug"}
	dog2 := Dogs{"Sweetie", "White", 6, 32, "Jerman Shepherd"}
	dog3 := Dogs{"Pluto", "Brown", 2, 19, "Bull Dog"}

	dog1.printAge()
	dog2.printBreed()
	dog3.printName()
}

func (d Dogs) printAge() {
	fmt.Println("Age of Dog is : ", d.Age)
}

func (d Dogs) printName() {
	fmt.Println("Name of Dog is : ", d.Name)
}

func (d Dogs) printBreed() {
	fmt.Println("Breed of the Dog is : ", d.Breed)
}
