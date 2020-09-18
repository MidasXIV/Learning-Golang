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
	fmt.Println("b &^ c = " + strconv.FormatInt(b&^c, 2)) // NOR
	fmt.Println("b << 3 = " + strconv.FormatInt(b<<3, 2)) // left shift
	fmt.Println("b >> 3 = " + strconv.FormatInt(b>>3, 2)) // right shift

	// dealing with complex numbers
	var n complex64 = 45 + 34i
	var m complex128 = complex(45, 65)
	fmt.Printf("%v , %T\n", real(n), real(n))
	fmt.Printf("%v , %T\n", imag(n), imag(n))
	fmt.Printf("%v , %T\n", n, n)
	fmt.Printf("%v , %T\n", imag(m), imag(m))
	fmt.Printf("%v , %T\n", real(m), real(m))
	fmt.Printf("%v , %T\n", m, m)

	// dealing with strings
	// strings in Go are any UTF8 characters

	s := "strings"
	s2 := ",another strings"
	by := []byte(s)
	fmt.Printf("%v , %T\n", s, s)       // string
	fmt.Printf("%v , %T\n", s+s2, s+s2) // string
	fmt.Printf("%v , %T\n", by, by)     // []unit8
	// s[2] = t // strings are immutable
	fmt.Printf("%v , %T\n", s[2], s[2]) //unit8

	// dealing with rune
	// rune in Go are any UTF32 characters

	r := 'a'  // rune
	r2 := "a" // string

	fmt.Printf("%v , %T\n", r, r)   // 97 , int32
	fmt.Printf("%v , %T\n", r2, r2) // a , string
}
