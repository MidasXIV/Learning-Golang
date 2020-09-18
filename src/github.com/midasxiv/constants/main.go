package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning About Constants")
	const myConst int = 32 // used when variable is not being exported      (camel Case)
	const MyConst int = 45 // start with capital letter when being exported (pascal Case)

	// const my2Const float64 = math.Sin(45)
	// cannot declare something a constant that is computed at run time

	fmt.Printf("%v , %T\n", myConst, myConst)

	var (
		a       = 22
		b int16 = 33
	)

	fmt.Printf("%v , %T\n", a, a)
	fmt.Printf("%v , %T\n", b, b)
	// fmt.Printf("%v , %T\n", a+b, a+b) // fails as a and b are different types

	const (
		c       = 22
		d int16 = 33
	)

	fmt.Printf("%v , %T\n", c, c) // int
	fmt.Printf("%v , %T\n", d, d) // int16
	// Implicit conversions when working with constants
	fmt.Printf("%v , %T\n", c+d, c+d) //int16
	// works here because c is a constant and when program is complied
	// each instance of d is replaced with 33 hence type is not considered

	// Learning about "iota"
	// enumerated constants
	const (
		e = iota
		f = iota
		g = iota
	)
	fmt.Printf("%v , %T\n", e, e) // 0
	fmt.Printf("%v , %T\n", f, f) // 1
	fmt.Printf("%v , %T\n", g, g) // 2

	// Go infers patterns
	// iota is scoped to a constant block
	// meaning if iota is used in different constant block iota will be 0 again
	const (
		e1 = iota
		f2
		g2
	)
	fmt.Printf("%v , %T\n", e1, e1) // 0
	fmt.Printf("%v , %T\n", f2, f2) // 1
	fmt.Printf("%v , %T\n", g2, g2) // 2

	// USE of enumerated constants
	const (
		isAdmin = 1 << iota
		isHQ
		isManager
		isProgrammer
		isHR
		isPE
	)

	var management byte = isAdmin | isHR | isHQ | isManager
	var staff byte = isProgrammer | isPE

	fmt.Printf("%b\n", management)
	fmt.Printf("%b\n", staff)

	fmt.Printf("Is Admin a management position ? %v\n", isAdmin&management == isAdmin)
	fmt.Printf("Is PE a management position ? %v\n", isPE&management == isPE)
	fmt.Printf("Is Manager a staff position ? %v\n", isManager&staff == isManager)
}
