package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Learning Primitive Datatypes")

	var a bool = true
	fmt.Printf("%v , %T\n", a, a)

	// bit operators
	var (
		b int64 = 10
		c int64 = 3
	)

	fmt.Println("b = " + strconv.FormatInt(b, 2))
	fmt.Println("c = " + strconv.FormatInt(c, 2))
	fmt.Println("b & c = " + strconv.FormatInt(b&c, 2))   // AND
	fmt.Println("b | c = " + strconv.FormatInt(b|c, 2))   // OR
	fmt.Println("b ^ c = " + strconv.FormatInt(b^c, 2))   // XOR
	fmt.Println("b &^ c = " + strconv.FormatInt(b&^c, 2)) //
	fmt.Println("b << 3 = " + strconv.FormatInt(b<<3, 2)) // left shift
	fmt.Println("b >> 3 = " + strconv.FormatInt(b>>3, 2)) // right shift
}
