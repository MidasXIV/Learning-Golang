package main

import (
	"fmt"
)

// if we start with capital letter it needs to be exported
// from the package. if small then internal to the package
// even the fields names need to be in uppercase

// if you can this to "doctor" warning goes away

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	fmt.Println("Learning Structures in go")
	aDoctor := Doctor{
		number:    3,
		actorName: "Mxiv",
		companions: []string{
			"mxiv1", "mxiv2",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions)

	//without the properties
	// positional syntax the options passed are passed to the positions
	// not recommended
	// someone may change the order of the  properties and it'll be mismatched
	anotherDoctor := Doctor{
		3,
		"Mxiv2",
		[]string{
			"mxiv12", "mxiv22",
		},
	}

	fmt.Println(anotherDoctor)

	//anon structure
	aDoc := struct{ name string }{name: "Mxiv"}
	fmt.Println(aDoc)

	anotherDoc := aDoc // deep cloned/ a copy is made
	anotherDoc.name = "Mxiv Clone"

	pointerDoc := &aDoc // pointer
	pointerDoc.name = "Mxiv Point"

	fmt.Println(aDoc)
	fmt.Println(anotherDoc)

}
