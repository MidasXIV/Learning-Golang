package main

import "fmt"

func main() {
	fmt.Println("Learning Variables")

	// Type                       -> Zero Value
	// bool                       -> false
	// all numeric types          -> 0
	// single character           -> 	0
	// string                     -> ""
	// array                      -> array of proper length where all elements have their zero value
	// slice                      -> empty slice with length 0 and capacity 0
	// map                        -> empty map
	// struct                     -> struct where all fields have their zero value
	// pointers & all other types -> nil

	var firstName, age = "Mark", 57 // must initialize all
	var (                           // use parens to spread over multiple lines
		lastName string = "Twain"
		books    int    // not initialized, so uses zero value
	)

	// all variables declared must be used else go gives an error
	fmt.Printf("%v %v | Age - %v | Books - %v", firstName, lastName, age, books)

	// no private variables
}
