package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning switch in Golang")
	i := 1
	fmt.Printf("%v, %T\n", i, i)

	//case 1, no implict fallthrough, Golang has implict breaks after each case
	// use fallthrough keyword to control flow
	// use break keyword to exit out of condition quickly
	switch i {
	case 1:
		fmt.Println("i is 1")
		fallthrough
		// prints out this and next case when case is 1
	case 2:
		fmt.Println("i is 2")
	default:
		fmt.Println("i is not 1 or 2")
	}

	// case 2 grouping cases
	// you can initilize a variable and then provide it as the "tag"
	switch tag := i + 5; tag {
	case 1, 3, 5, 7, 9:
		fmt.Printf("tag %v is odd\n", tag)
	case 2, 4, 6, 8:
		fmt.Printf("tag %v is even\n", tag)
	default:
		fmt.Printf("tag %v is not less than 10\n", tag)
	}

	// case 3 providing ranges, tag-less syntax
	// overlapping ranges, without providing a tag
	i = 15
	switch {
	case i <= 10:
		fmt.Println("i is <= 10")
	case i <= 20:
		fmt.Println("i is <= 20")
	default:
		fmt.Println("i > 20")
	}

	// case 4 - using interfaces
	var k interface{} = 15.0
	switch k.(type) {
	case int:
		fmt.Println("k is of type int")
	case float64:
		fmt.Println("k is of type float64")
	case string:
		fmt.Println("k is of type string")
	default:
		fmt.Println("k is of unknow type")
	}
}
