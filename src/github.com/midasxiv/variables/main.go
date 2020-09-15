package main

import "fmt"

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
}
