package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	Speed  float32
	CanFly bool
}

func main() {
	fmt.Println("Learning Embedding in Golang")
	// Go does not have inheritance
	// hence we use composistion if Bird needs to have properties
	// of Animal then compose Animal in type Bird

	aBird := Bird{
		// Name:   "Emu",             // cannot be used like this
		// Origin: "Australia",
		Speed:  48,
		CanFly: false,
	}

	aBird.Name = "Emu"
	aBird.Origin = "Australia"

	fmt.Println(aBird)

	// another way of declaring
	anotherBird := Bird{
		Animal: Animal{Name: "Emu", Origin: "Australia"},
		Speed:  48,
		CanFly: false,
	}

	fmt.Println(anotherBird)

	// here Bird "has-a" animal but Animal and Bird cannot be used Interchangiably
}
