package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Learning Variables")

	// var<keyword> <variable-name> <variable-type>
	var i int

	// initializing a variable
	i = 42
	i = 27 // variables can be reassigned

	// declaring a variable and assigning a value with type
	var j float32 = 32

	// declaring a variable and assigning a value
	// but let Go infer the type
	k := 34.0
	// Go infers interger to be of type int or float64

	fmt.Println(i, j, k)

	// print using formatted string
	// %v is used to print value
	// %T is used to print type
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)

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
	fmt.Printf("%v %v | Age - %v | Books - %v\n", firstName, lastName, age, books)

	// Go does not support private variables
	// instead decalred variables in package can be treated as private

	// converting variables from one type to another
	var m int = 36
	var n float64
	n = float64(m)

	var o string
	// converts int to ascii string
	o = strconv.Itoa(m)

	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", o, o)

}
